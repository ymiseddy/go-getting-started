package ephemeral

import "github.com/ymiseddy/go-getting-started/internal/data"

type EphemeralService struct {
	database *data.Database
}

func NewEphemeralService(database *data.Database) *EphemeralService {
	{
		return &EphemeralService{database: database}
	}
}
