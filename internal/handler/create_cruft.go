package handler

import (
	"context"

	serviceProto "github.com/chrusty/go-grpc-service/pkg/api/v1/go"
)

// CreateCruft adds cruft to storage:
func (h *Handler) CreateCruft(ctx context.Context, req *serviceProto.CreateCruftRequest) (*serviceProto.CreateCruftResponse, error) {
	h.logger.Tracef("Handling a CreateCruft() request...")

	// Store the cruft:
	cruftID, err := h.storer.CreateCruft(req.Cruft)
	if err != nil {
		h.logger.Warnf("Unable to store cruft: %v", err)
		return nil, err
	}

	// Prepare a response:
	return &serviceProto.CreateCruftResponse{
		StatusCode: serviceProto.StatusCode_OK,
		CruftId:    cruftID,
	}, nil
}
