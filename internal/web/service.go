package web

import (
	"github.com/ymiseddy/go-getting-started/internal/data"
	"github.com/ymiseddy/go-getting-started/internal/request"
)

type WebServerConfig struct {
	Host string `env:"WEB_HOST,default=0.0.0.0"`
	Port int    `env:"WEB_PORT,default=8080"`
}

type WebServer struct {
	config              WebServerConfig
	database            *data.Database
	requestInfoProvider request.RequestInfoProvider
}

func NewWebServer(config WebServerConfig, database *data.Database, requestInfoProvider request.RequestInfoProvider) *WebServer {
	return &WebServer{
		config:              config,
		database:            database,
		requestInfoProvider: requestInfoProvider,
	}
}
