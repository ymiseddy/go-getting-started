package data

type DatabaseConfig struct {
	DataSourceName string `env:"DB_DSN,required"`
}

type Database struct {
	config *DatabaseConfig
}

func NewDatabase(config *DatabaseConfig) *Database {
	return &Database{config: config}
}
