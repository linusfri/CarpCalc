package db

import (
	"fmt"
	"os"

	user "github.com/linusfri/calc-api/models/user"

	"github.com/linusfri/calc-api/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Gorm *gorm.DB
}

func Connect() (*DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PW"),
		os.Getenv("DB_SERVICE"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	connection, err := gorm.Open(mysql.Open(connectionString))

	return &DB{connection}, err
}

func Migrate() {
	db, err := Connect()

	if err != nil {
		helper.HandleErr(err)
		return
	}

	db.Gorm.AutoMigrate(&user.User{})
}
