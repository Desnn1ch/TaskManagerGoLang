package config

import (
	"TaskManagerGoLang/models"
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

const (
	kConfigName string = "config"
	kConfigType string = "yaml"
	kConfigPath string = "."
)

// init инициализирует viper один раз для всех функций
func init() {
	viper.SetConfigName(kConfigName)
	viper.SetConfigType(kConfigType)
	viper.AddConfigPath(kConfigPath)
}

func LoadDataBaseConfig() (models.DataBase, error) {
	var dbConfig models.DataBase

	if err := viper.ReadInConfig(); err != nil {
		return dbConfig, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := viper.UnmarshalKey("database", &dbConfig); err != nil {
		return dbConfig, fmt.Errorf("failed to unmarshal database config: %w", err)
	}

	return dbConfig, nil
}

func LoadSecretKeyConfig() (string, error) {
	var jwtConfig models.JWT

	if err := viper.ReadInConfig(); err != nil {
		return "", fmt.Errorf("failed to read config file: %w", err)
	}

	if err := viper.UnmarshalKey("JWT", &jwtConfig); err != nil {
		return "", fmt.Errorf("failed to unmarshal JWT config: %w", err)
	}

	if jwtConfig.Secret_key == "" {
		return "", errors.New("secret key is empty")
	}

	return jwtConfig.Secret_key, nil
}

func LoadRedisConfig() (models.Redis, error) {
	var redisConfig models.Redis

	if err := viper.ReadInConfig(); err != nil {
		return redisConfig, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := viper.UnmarshalKey("redis", &redisConfig); err != nil {
		return redisConfig, fmt.Errorf("failed to unmarshal redis config: %w", err)
	}

	return redisConfig, nil
}
