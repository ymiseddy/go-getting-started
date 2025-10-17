package main

import (
	"fmt"
	"github.com/ymiseddy/go-getting-started/internal/config"
)

func main() {
	appConfig, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(appConfig.Message)
}
