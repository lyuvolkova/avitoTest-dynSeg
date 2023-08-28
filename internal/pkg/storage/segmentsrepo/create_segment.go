package segmentsrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
)

const createSegmentQuery = `
INSERT INTO segments (slug)
VALUES ($1)
`

var ErrDuplicate = errors.New("duplicate")

func (r *repository) CreateSegment(ctx context.Context, slug string) error {
	res, err := r.db.Exec(ctx, createSegmentQuery, slug)
	var pgErr *pgconn.PgError
	switch {
	case errors.Is(err, nil): // pass
	case errors.As(err, &pgErr):
		if pgErr.Code == "23505" {
			return ErrDuplicate
		}

		return err
	default:
		return err
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("No rows inserted")
	}

	return nil
}
