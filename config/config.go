package config

import (
	"TaskManagerGoLang/models"

	"github.com/spf13/viper"
)

func LoadConfig() (models.DataBase, error) {
	var dbConfig models.DataBase

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
