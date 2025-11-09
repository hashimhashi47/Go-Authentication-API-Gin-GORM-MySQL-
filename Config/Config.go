package config

import (
	model "authentication/Model"
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


//connecting with database(mysql)
func DBconfig() {
	DNS := "root:**********(127.0.0.1:3306)/Data"
	var err error
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal("ERROR OCCURED ON CONNECTION BETWEEN DATABASE")
	}

	fmt.Println("DATABASE CONNECTED SUCESSFULLY")

	DB.AutoMigrate(&model.User{})
}
