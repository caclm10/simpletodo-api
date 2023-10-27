package app

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func NewConfig(d ...string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config := viper.New()

	dir := filepath.Join(append([]string{wd}, d...)...)
	config.SetConfigFile(filepath.Join(dir, ".env"))

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	Config = config
}
