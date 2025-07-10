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

	var rawConfig struct {
		PropertySources []struct {
			Source map[string]string `json:"source"`
		} `json:"propertySources"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&rawConfig); err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, source := range rawConfig.PropertySources {
		for k, v := range source.Source {
			result[k] = v
		}
	}
	return result, nil
}


func init() {
	configUrl := os.Getenv("CONFIG_SERVER")
	ConfigSettings = &Config{}
	if configUrl != "" {
		for {
			config, err := fetchConfigFromServer(configUrl + "/users-api/default")
			log.Println(config)
			if err == nil {
				// Tenta pegar tanto DATABASE_URL quanto database.url
				ConfigSettings.DatabaseUrl = config[constants.DATABASE_URL]
				if ConfigSettings.DatabaseUrl == "" {
					ConfigSettings.DatabaseUrl = config["database.url"]
				}
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
