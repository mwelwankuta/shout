package database

import (
	"log"

	"github.com/mwelwankuta/shout/server/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DNS = "root:@tcp(127.0.0.1:3306)/gotut"

func Connect() {
	connection, err := gorm.Open(mysql.Open(DNS))

	if err != nil {
		panic("Could not connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
	log.Printf("Connected to the database")
}
