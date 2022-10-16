package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	name := os.Getenv("mysqluser")
	password := os.Getenv("mysqluserpassword")
	dsn := string(name) + ":" + string(password) + "@tcp(127.0.0.1:3306)/atm?charset=utf8mb4&parseTime=True&loc=Local"
	dbTable1, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the Users DB")
	}
	return dbTable1
}
