package ioc

import (
	"context"

	"github.com/ymiseddy/go-getting-started/internal/config"
	"github.com/ymiseddy/go-getting-started/internal/data"
	"github.com/ymiseddy/go-getting-started/internal/ephemeral"
	"github.com/ymiseddy/go-getting-started/internal/request"
	"github.com/ymiseddy/go-getting-started/internal/tenant"
	"github.com/ymiseddy/go-getting-started/internal/web"
)

type ServiceProvider struct {
	databaseConfig  *data.DatabaseConfig
	database        *data.Database
	webServerConfig *web.WebServerConfig
	webServer       *web.WebServer
	tenantServices  map[string]*tenant.TenantService
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{
		tenantServices: make(map[string]*tenant.TenantService),
	}
}

func (sp *ServiceProvider) GetDatabaseConfig() (*data.DatabaseConfig, error) {
	if sp.database == nil {
		databaseConfig := data.DatabaseConfig{}
		err := config.ReadConfigInto(&databaseConfig)
		if err != nil {
			return nil, err
		}
		sp.databaseConfig = &databaseConfig
	}
	return sp.databaseConfig, nil
}

func (sp *ServiceProvider) GetDatabase() (*data.Database, error) {
	if sp.database == nil {
		dbConfig, err := sp.GetDatabaseConfig()
		if err != nil {
			return nil, err
		}
		sp.database = data.NewDatabase(dbConfig)
	}
	return sp.database, nil
}

func (sp *ServiceProvider) GetWebServerConfig() (*web.WebServerConfig, error) {
	if sp.webServer == nil {
		webServerConfig := web.WebServerConfig{}
		err := config.ReadConfigInto(&webServerConfig)
		if err != nil {
			return nil, err
		}
		sp.webServerConfig = &webServerConfig
	}
	return sp.webServerConfig, nil
}

func (sp *ServiceProvider) GetWebServer() (*web.WebServer, error) {
	if sp.webServer == nil {
		wsConfig, err := sp.GetWebServerConfig()
		if err != nil {
			return nil, err
		}
		database, err := sp.GetDatabase()
		if err != nil {
			return nil, err
		}
		sp.webServer = web.NewWebServer(*wsConfig, database, sp)
	}
	return sp.webServer, nil
}

func (sp *ServiceProvider) GetTenantServiceFor(name string) (*tenant.TenantService, error) {
	service, exists := sp.tenantServices[name]
	if !exists {
		database, err := sp.GetDatabase()
		if err != nil {
			return nil, err
		}
		service = tenant.NewTenantService(name, database)
		sp.tenantServices[name] = service
	}
	return service, nil
}

func (sp *ServiceProvider) GetAnEphemeralService() (*ephemeral.EphemeralService, error) {
	database, err := sp.GetDatabase()
	if err != nil {
		return nil, err
	}
	service := ephemeral.NewEphemeralService(database)
	return service, nil
}

type requestInfoKeyType string

const requestInfoKey requestInfoKeyType = "requestInfo"

func (sp *ServiceProvider) GetRequestInfo(ctx context.Context) (*request.RequestInfo, context.Context, error) {
	reqInfo, ok := ctx.Value(requestInfoKey).(*request.RequestInfo)
	if !ok {
		reqInfo = &request.RequestInfo{}
		ctx = context.WithValue(ctx, requestInfoKey, reqInfo)
	}
	return reqInfo, ctx, nil
}
