package service

import "context"

type segmentsRepo interface {
	CreateSegment(ctx context.Context, slug string) error
	UpdateUserSegments(ctx context.Context, id int64, addSlugs []string, deleteSlugs []string) error
	DeleteSegment(ctx context.Context, slug string) error
	GetUserSegments(ctx context.Context, id int64) ([]string, error)
}
