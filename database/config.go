package database

import (
	"fmt"

	models "github.com/mwelwankuta/gotut/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DNS = "root:@tcp(127.0.0.1:3306)/gotut"

func InitializeDatabase() {
	var connection, err = gorm.Open(mysql.Open(DNS))

	if err != nil {
		panic("Could not connect to database")
	}
	DB = connection
	fmt.Println("Database Connected successfully")

	DB.AutoMigrate(models.User{})
	DB.AutoMigrate(models.Post{})
}
