package config

import (
	"os"
)

// Primaryc application configuration structure
type AppConfig struct {
	Message string
}

func ReadConfig() (*AppConfig, error) {
	message, found := os.LookupEnv("APP_MESSAGE")
	if !found {
		// Of course, we could return an error here instead.
		message = "Hello, World!"
	}
	return &AppConfig{
		Message: message,
	}, nil

}
