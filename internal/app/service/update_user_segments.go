package service

import (
	"context"
	"fmt"
	"github.com/lyuvolkova/avitoTest-dynSeg/internal/pkg/server"
)

func (a *app) UpdateUserSegments(ctx context.Context, req *server.UpdateUserSegmentsRequest) (*server.UpdateUserSegmentsResponse, error) {
	err := a.sp.UpdateUserSegments(ctx, req.UserID, req.AddSlug, req.DeleteSlug)
	if err != nil {
		return nil, fmt.Errorf("SP update segments: %w", err)
	}
	return &server.UpdateUserSegmentsResponse{Ok: true}, nil
}
