package app

import (
	"fmt"

	"github.com/caclm10/simpletodo-api/internal/config"
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
	Connection: config.Viper.GetString("DB_CONNECTION"),
	Host:       config.Viper.GetString("DB_HOST"),
	Port:       config.Viper.GetString("DB_PORT"),
	Database:   config.Viper.GetString("DB_DATABASE"),
	Username:   config.Viper.GetString("DB_USERNAME"),
	Password:   config.Viper.GetString("DB_PASSWORD"),
}

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconfig.Username, dbconfig.Password, dbconfig.Host, dbconfig.Port, dbconfig.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}
