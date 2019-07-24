package handler

import (
	"context"

	storage "github.com/chrusty/go-grpc-service/internal/storage"
	serviceProto "github.com/chrusty/go-grpc-service/pkg/api/v1/go"
)

// ReadCruft retrieves cruft from storage:
func (h *Handler) ReadCruft(ctx context.Context, req *serviceProto.ReadCruftRequest) (*serviceProto.ReadCruftResponse, error) {
	h.logger.Tracef("Handling a ReadCruft() request...")

	// Read the cruft:
	cruft, err := h.storer.ReadCruft(req.CruftId)
	if err != nil {
		switch err {
		case storage.ErrNotFound:
			return &serviceProto.ReadCruftResponse{
				StatusCode: serviceProto.StatusCode_CRUFT_NOT_FOUND,
			}, nil
		default:
			return nil, err
		}
	}

	// Prepare a response:
	return &serviceProto.ReadCruftResponse{
		StatusCode: serviceProto.StatusCode_OK,
		Cruft:      cruft,
	}, nil
}
