package laptop

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	pb_laptop "github.com/yangwawa0323/pcbook/pb/laptop/v1"
	pb_user "github.com/yangwawa0323/pcbook/pb/user/v1"

	"github.com/jinzhu/copier"
	"github.com/yangwawa0323/pcbook/sql"
	"github.com/yangwawa0323/pcbook/utils"
	"gorm.io/gorm"
)

var ErrAlreadyExists = errors.New("record already exists")
var out = utils.NewDebugOutput()

// LaptopStore is an interface to store laptop
type LaptopStore interface {
	Save(context.Context, *pb_laptop.Laptop) error
	Find(ctx context.Context, id string) (*pb_laptop.Laptop, error)
}

// DbLaptopStore stores laptop in db
type DbLaptopStore struct {
	DB *gorm.DB
}

func NewDbLaptopStore() (dbLaptopStore *DbLaptopStore) {

	db, err := sql.InitDB()
	if err != nil {
		log.Fatal(out.Panic("cannot connect to MySQL database: %v", err))
	}

	dbLaptopStore = &DbLaptopStore{
		DB: db,
	}
	err = dbLaptopStore.Migrate()
	if err != nil {
		log.Fatal(out.Panic("cannot migrate to MySQL database: %v", err))
	}
	return
}

func (store *DbLaptopStore) Migrate() error {
	return store.DB.AutoMigrate(
		&pb_laptop.LaptopORM{},
		&pb_laptop.CpuORM{},
		&pb_laptop.GpuORM{},
		&pb_laptop.ScreenORM{},
		&pb_laptop.KeyboardORM{},
		&pb_laptop.StorageORM{},
		&pb_user.UserORM{},
		&pb_user.EmailORM{},
		&pb_user.CreditCardORM{},
	)
}

func (store *DbLaptopStore) Save(ctx context.Context, laptop *pb_laptop.Laptop) error {
	laptopOrm, err := laptop.ToORM(ctx)
	if err != nil {
		log.Fatal(out.Panic("cannot convert protobuff to ORM: %v", err))
		return err
	}
	log.Printf("%#v", laptopOrm)
	result := store.DB.Create(&laptopOrm)
	if result.Error != nil {
		log.Fatal(out.Panic("cannont save laptopOrm to database: %v", err))
		return err
	}
	return nil
}

func (store *DbLaptopStore) Find(ctx context.Context, id string) (*pb_laptop.Laptop, error) {
	// laptopORM := pb_laptop.LaptopORM{
	// 	Id: id,
	// }

	var laptopORM pb_laptop.LaptopORM
	store.DB.Where(" id = ? ", id).First(&laptopORM)

	log.Print(out.Debug("GORM: %#v\n", laptopORM))
	laptop, err := laptopORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &laptop, nil
}

// InMemoryLaptopStore stores laptop in memory
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb_laptop.Laptop
}

// NewInMemoryLaptopStore returns a new InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb_laptop.Laptop),
	}
}

func (store *InMemoryLaptopStore) Save(ctx context.Context, laptop *pb_laptop.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	// deep copy
	other := &pb_laptop.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}

	store.data[other.Id] = other
	return nil
}

func (store *InMemoryLaptopStore) Find(ctx context.Context, id string,
) (*pb_laptop.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	// deep copy
	other := &pb_laptop.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	return other, nil

}
