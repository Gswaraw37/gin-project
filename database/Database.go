package database

import (
	"fmt"
	dbconfig "gin-project/config/db_config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	if dbconfig.DB_DRIVER == "pgsql" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbconfig.DB_HOST, dbconfig.DB_USER, dbconfig.DB_PASSWORD, dbconfig.DB_NAME, dbconfig.DB_PORT)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if dbconfig.DB_DRIVER == "mysql" {
		dsnMysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconfig.DB_USER, dbconfig.DB_PASSWORD, dbconfig.DB_HOST, dbconfig.DB_PORT, dbconfig.DB_NAME)
		DB, err = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{})
	}

	if err != nil {
		panic("Can't connect to database")
	}

	log.Println("Connected to database")
}
