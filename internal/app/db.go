package app

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var dbconfig = struct {
	Connection string
	Host       string
	Port       string
	Database   string
	Username   string
	Password   string
}{
	Connection: Config.GetString("DB_CONNECTION"),
	Host:       Config.GetString("DB_HOST"),
	Port:       Config.GetString("DB_PORT"),
	Database:   Config.GetString("DB_DATABASE"),
	Username:   Config.GetString("DB_USERNAME"),
	Password:   Config.GetString("DB_PASSWORD"),
}

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconfig.Username, dbconfig.Password, dbconfig.Host, dbconfig.Port, dbconfig.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}
