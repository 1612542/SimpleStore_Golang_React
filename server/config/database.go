package config

import (
	"log"

	"github.com/spf13/viper"
)

type MysqlConfig struct {
	DbDriver string
	DbHost   string
	DbPort   string
	DbName   string
	UserName string
	Password string
}

func MySQLConfig() MysqlConfig {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	return MysqlConfig{
		DbDriver: viper.GetString("DB_DRIVER"),
		DbHost:   viper.GetString("DB_HOST"),
		DbPort:   viper.GetString("DB_PORT"),
		DbName:   viper.GetString("DB_NAME"),
		UserName: viper.GetString("USERNAME"),
		Password: viper.GetString("PASSWORD"),
	}
}
