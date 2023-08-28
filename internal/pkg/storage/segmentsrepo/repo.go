package segmentsrepo

import "github.com/lyuvolkova/avitoTest-dynSeg/internal/pkg/storage"

type repository struct {
	db storage.Database
}

func New(db storage.Database) *repository {
	return &repository{
		db: db,
	}
}
