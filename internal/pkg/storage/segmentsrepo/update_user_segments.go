package segmentsrepo

import (
	"context"
	"github.com/lib/pq"
)

const addSegmentQuery = `
INSERT INTO segment_users
SELECT "segmentID", $1 AS userID FROM segments WHERE slug=any($2) ON CONFLICT DO NOTHING 
`
const deleteUserSegmentQuery = `
DELETE FROM segment_users WHERE "userID" = $1 AND "segmentID" IN
                                                  (SELECT "segmentID" FROM segments WHERE slug=any($2))
`

func (r *repository) UpdateUserSegments(ctx context.Context, id int64, addSlugs []string, deleteSlugs []string) error {
	if len(addSlugs) != 0 {
		_, err := r.db.Exec(ctx, addSegmentQuery, id, pq.Array(addSlugs))
		if err != nil {
			return err
		}
	}
	if len(deleteSlugs) != 0 {
		_, err := r.db.Exec(ctx, deleteUserSegmentQuery, id, pq.Array(deleteSlugs))
		if err != nil {
			return err
		}

	}
	return nil
}
