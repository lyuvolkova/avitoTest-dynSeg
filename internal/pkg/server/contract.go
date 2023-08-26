package server

import "context"

type app interface {
	CreateSegment(context.Context, *CreateSegmentRequest) (*CreateSegmentResponse, error)
	DeleteSegment(context.Context, *DeleteSegmentRequest) (*DeleteSegmentResponse, error)
	GetUserSegments(context.Context, *GetUserSegmentsRequest) (*GetUserSegmentsResponse, error)
	UpdateUserSegments(context.Context, *UpdateUserSegmentsRequest) (*UpdateUserSegmentsResponse, error)
}
