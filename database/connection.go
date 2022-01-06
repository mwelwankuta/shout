package database

import (
	"log"

	"github.com/mwelwankuta/shout/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DNS = "root:@tcp(127.0.0.1:3306)/gotut"

func Connect() error {
	connection, err := gorm.Open(mysql.Open(DNS))

	DB = connection

	connection.AutoMigrate(&models.User{})
	log.Printf("Connected to the database")
	return err
}
