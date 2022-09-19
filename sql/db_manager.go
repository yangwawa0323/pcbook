package sql

import (
	"fmt"
	"log"
	"os"

	"github.com/yangwawa0323/pcbook/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var out = utils.NewDebugOutput()

func GetMySqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"redhat",
		"localhost:3306",
		"testing",
	)
}

func InitDB() (*gorm.DB, error) {
	infoLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	// db, err := gorm.Open(mysql.Open(GetMySqlDSN()), &gorm.Config{Logger: infoLogger})
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{Logger: infoLogger})
	if err != nil {
		log.Fatal(out.Panic("cannot connect to MySQL database: %v", err))
		return nil, err
	}
	return db, err
}
