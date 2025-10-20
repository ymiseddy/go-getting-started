package request

import (
	"context"
	"math/rand"
)

type RequestInfo struct {
	ID       int64
	username string
}

type RequestInfoProvider interface {
	GetRequestInfo(ctx context.Context) (*RequestInfo, context.Context, error)
}

func NewRequestInfo() *RequestInfo {
	id := generateUniqueID()
	return &RequestInfo{ID: id}
}

func generateUniqueID() int64 {
	// Generate random unique ID
	id := rand.Int63()
	return id
}
