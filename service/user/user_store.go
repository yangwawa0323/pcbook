package user

import (
	"fmt"
	"log"


	pb_user "github.com/yangwawa0323/pcbook/pb/user/v1"

	"github.com/yangwawa0323/pcbook/sql"
	"github.com/yangwawa0323/pcbook/utils"
	"gorm.io/gorm"
)

var out = utils.DebugOutput{}

type UserStore interface {
	Save(*pb_user.UserORM) error
	Find(id string) (*pb_user.UserORM, error)
}

type UserDBStore struct {
	DB *gorm.DB
}

func GetMySqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"redhat",
		"localhost:3306",
		"testing",
	)
}

func NewUserDBStore() *UserDBStore {

	db, err := sql.InitDB()
	if err != nil {
		log.Fatal(out.Panic("cannot connect to MySQL database: %v", err))
	}

	userDBStore := &UserDBStore{
		DB: db,
	}
	err = userDBStore.Migrate()
	if err != nil {
		log.Fatal(out.Panic("cannot migrate to MySQL database: %v", err))
	}
	return userDBStore
}

func (store *UserDBStore) Migrate() error {
	return store.DB.AutoMigrate(
		&pb_user.UserORM{},
		&pb_user.EmailORM{},
		&pb_user.AddressORM{},
	)
}

func (store *UserDBStore) Save(user *pb_user.UserORM) error {
	return store.DB.Create(user).Error
}

func (store *UserDBStore) Find(id string) (*pb_user.UserORM, error) {
	var user *pb_user.UserORM
	err := store.DB.Joins("Email").First(&user, "id = ? ", id).Error
	return user, err
}
