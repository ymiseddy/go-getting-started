package main

import (
	"fmt"
	"github.com/ymiseddy/go-getting-started/internal/config"
)

func main() {
	appConfig := config.AppConfig{}
	err := config.ReadConfigInto(&appConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(appConfig.Message)

	fmt.Printf("Here's our config: %+v\n", appConfig)
}
