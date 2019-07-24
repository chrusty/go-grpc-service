package handler

import (
	storage "github.com/chrusty/go-grpc-service/internal/storage"
	sharedInterfaces "github.com/chrusty/go-grpc-service/pkg/interfaces"
)

// Handler implements the service proto:
type Handler struct {
	logger sharedInterfaces.Logger
	storer storage.Storer
}

// New returns a new handler instantiated with a provided Logger:
func New(logger sharedInterfaces.Logger, storer storage.Storer) *Handler {

	logger.Infof("Created a new Handler")

	return &Handler{
		logger: logger,
		storer: storer,
	}
}
