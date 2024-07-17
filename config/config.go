package config

import (
	"github.com/spf13/viper"
)

type DataBase struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
	SSLMode  string
}

func LoadConfig() (DataBase, error) {
	var dbConfig DataBase

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return dbConfig, err
	}

	err := viper.UnmarshalKey("database", &dbConfig)
	if err != nil {
		return dbConfig, err
	}

	return dbConfig, nil
}
