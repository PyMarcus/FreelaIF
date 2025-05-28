package config

import (
	"log"

	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/constants"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseUrl string
}

var ConfigSettings *Config

func init() {
	viper.SetConfigFile(constants.ENV_PATH)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Missing .env: %v", err)
	}

	ConfigSettings = &Config{}

	ConfigSettings.DatabaseUrl = viper.GetString(constants.DATABASE_URL)
}
