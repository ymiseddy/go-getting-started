package tenant

import "github.com/ymiseddy/go-getting-started/internal/data"

type TenantService struct {
	name     string
	database *data.Database
}

func NewTenantService(name string, database *data.Database) *TenantService {
	return &TenantService{name: name, database: database}
}
