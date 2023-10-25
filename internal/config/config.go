package config

import (
	"os"

	"github.com/spf13/viper"
)

var Viper *viper.Viper

func NewViper() {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config := viper.New()

	config.SetConfigFile(".env")
	config.AddConfigPath(root)

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	Viper = config
}
