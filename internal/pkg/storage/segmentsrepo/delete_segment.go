package segmentsrepo

import (
	"context"

	"github.com/jackc/pgx/v4"
)

const deleteSegmentQuery = `
DELETE FROM segment_users WHERE "segmentID" IN
                                (SELECT "segmentID" FROM segments WHERE slug = $1);
DELETE FROM segments WHERE slug = $1
`

func (r *repository) DeleteSegment(ctx context.Context, slug string) error {
	return r.db.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, deleteSegmentQuery, slug)
		if err != nil {
			return err
		}
		return nil
	})
}
