package config

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPAddr string `envconfig:"HTTP_ADDR"`
	PgURL    string `envconfig:"PG_URL"`
	LogLevel string `envconfig:"LOG_LEVEL"`
}

var (
	config Config
	once   sync.Once
)

// Get reads config from environment. Once.
func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}

		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Configuration:", string(configBytes))
	})

	return &config
}
