package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/lyuvolkova/avitoTest-dynSeg/internal/pkg/server"
	"github.com/lyuvolkova/avitoTest-dynSeg/internal/pkg/storage/segmentsrepo"
)

func (a *app) CreateSegment(ctx context.Context, req *server.CreateSegmentRequest) (*server.CreateSegmentResponse, error) {
	if req.Slug == "" {
		return nil, fmt.Errorf("%w: Incorrect slug", server.ErrBadRequest)
	}
	err := a.sp.CreateSegment(ctx, req.Slug)
	if err != nil {
		if errors.Is(err, segmentsrepo.ErrDuplicate) {
			return nil, fmt.Errorf("%w: segment exists", server.ErrConflict)
		}
		return nil, fmt.Errorf("SP create segment: %w", err)
	}

	return &server.CreateSegmentResponse{Ok: true}, nil
}
