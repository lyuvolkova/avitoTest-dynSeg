package service

import (
	"context"

	"github.com/lyuvolkova/avitoTest-dynSeg/internal/pkg/server"
)

func (a *app) DeleteSegment(context.Context, *server.DeleteSegmentRequest) (*server.DeleteSegmentResponse, error) {
	return nil, server.ErrNotImplemented
}
