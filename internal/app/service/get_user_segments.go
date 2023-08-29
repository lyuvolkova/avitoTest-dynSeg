package service

import (
	"context"
	"fmt"

	"github.com/lyuvolkova/avitoTest-dynSeg/internal/pkg/server"
)

func (a *app) GetUserSegments(ctx context.Context, req *server.GetUserSegmentsRequest) (*server.GetUserSegmentsResponse, error) {
	slugs, err := a.sp.GetUserSegments(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("SP get user segments: %w", err)
	}
	return &server.GetUserSegmentsResponse{Slugs: slugs}, nil
}
