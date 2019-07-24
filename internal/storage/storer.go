package storage

import (
	serviceProto "github.com/chrusty/go-grpc-service/pkg/api/v1/go"
)

// Storer deals with storing Cruft:
type Storer interface {
	CreateCruft(cruft *serviceProto.Cruft) (string, error)
	ReadCruft(cruftID string) (*serviceProto.Cruft, error)
}
