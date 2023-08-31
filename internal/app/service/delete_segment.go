package service

import (
	"context"
	"fmt"

	"github.com/lyuvolkova/avitoTest-dynSeg/internal/pkg/server"
)

func (a *app) DeleteSegment(ctx context.Context, req *server.DeleteSegmentRequest) (*server.DeleteSegmentResponse, error) {
	if req.Slug == "" {
		return nil, fmt.Errorf("%w: Incorrect slug", server.ErrBadRequest)
	}
	err := a.sp.DeleteSegment(ctx, req.Slug)
	if err != nil {
		return nil, fmt.Errorf("SP delete segment: %w", err)
	}
	return &server.DeleteSegmentResponse{Ok: true}, nil
}
