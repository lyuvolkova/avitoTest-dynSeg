package service

import (
	"context"

	"github.com/lyuvolkova/avitoTest-dynSeg/internal/pkg/server"
)

func (a *app) UpdateUserSegments(context.Context, *server.UpdateUserSegmentsRequest) (*server.UpdateUserSegmentsResponse, error) {
	return nil, server.ErrNotImplemented
}
