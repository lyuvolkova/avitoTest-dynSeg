package service

import (
	"context"

	"github.com/lyuvolkova/avitoTest-dynSeg/internal/pkg/server"
)

func (a *app) GetUserSegments(context.Context, *server.GetUserSegmentsRequest) (*server.GetUserSegmentsResponse, error) {
	return nil, server.ErrNotImplemented
}
