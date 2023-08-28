package service

import "context"

type segmentsRepo interface {
	CreateSegment(ctx context.Context, slug string) error
}
