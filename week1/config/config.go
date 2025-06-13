package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DataBase struct {
		Dsn          string
		MaxOpenConns int
		MaxIdleConns int
	}

	Redis struct {
		Cluster  bool
		Addrs    []string
		Password string
		PoolSize int
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %v", err)
	}

	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Error unmarshalling config, %v", err)
	}

	initDB()
	initRedis()
}
