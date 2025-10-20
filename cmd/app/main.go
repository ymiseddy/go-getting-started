package main

import (
	"fmt"

	"github.com/ymiseddy/go-getting-started/internal/ioc"
)

func main() {
	provider := ioc.NewServiceProvider()
	webServer, err := provider.GetWebServer()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Web server %+v\n", webServer)
}
