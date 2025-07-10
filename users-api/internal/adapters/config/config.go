package config

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/constants"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseUrl string
}

var ConfigSettings *Config

func fetchConfigFromServer(url string) (map[string]string, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, err
	}
	var config map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

func init() {
	configUrl := os.Getenv("CONFIG_URL")
	ConfigSettings = &Config{}
	if configUrl != "" {
		for {
			config, err := fetchConfigFromServer(configUrl)
			if err == nil {
				ConfigSettings.DatabaseUrl = config[constants.DATABASE_URL]
				if ConfigSettings.DatabaseUrl != "" {
					break
				}
			}
			log.Println("Aguardando config-server...")
			time.Sleep(3 * time.Second)
		}
	} else {
		viper.SetConfigFile(constants.ENV_PATH)
		viper.SetConfigType("env")
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Missing .env: %v", err)
		}
		ConfigSettings.DatabaseUrl = viper.GetString(constants.DATABASE_URL)
	}
}
