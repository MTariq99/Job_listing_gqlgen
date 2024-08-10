package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	MongoURL   string `mapstructure:"MONGO_URL"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	JwtSecret  string `mapstructure:"JWT_SECRET"`
}

var Cfg Config

func LoadConfig() error {
	viper.AddConfigPath("./")
	if os.Getenv("ENV") == "prod" {
		viper.SetConfigName(".env.prod")
	} else if os.Getenv("ENV") == "staging" {
		viper.SetConfigName(".env.stag")
	} else {
		viper.SetConfigName(".env.dev")
	}
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		return err
	}
	return nil
}
