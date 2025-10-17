package main

import (
	"fmt"
	"github.com/ymiseddy/go-getting-started/internal/config"
)

func main() {

	appConfig := config.AppConfig{}
	databaseConfig := config.DatabaseConfig{}
	webServerConfig := config.WebServerConfig{}

	err := config.ReadConfigInto(&appConfig)
	if err != nil {
		panic(err)
	}

	err = config.ReadConfigInto(&databaseConfig)
	if err != nil {
		panic(err)
	}

	err = config.ReadConfigInto(&webServerConfig)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Application config: %+v\n", appConfig)
	fmt.Printf("Database Config: %+v\n", databaseConfig)
	fmt.Printf("Web Server Config: %+v\n", webServerConfig)

	println(appConfig.Message)

}
