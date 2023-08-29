package segmentsrepo

import (
	"context"
)

const getUserSegmentsQuery = `
SELECT slug FROM segments
            WHERE "segmentID" IN
                  (SELECT "segmentID" FROM segment_users WHERE "userID" = $1)
`

func (r *repository) GetUserSegments(ctx context.Context, id int64) ([]string, error) {
	rows, err := r.db.Query(ctx, getUserSegmentsQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var slugs []string
	for rows.Next() {
		var slug string
		err = rows.Scan(&slug)
		if err != nil {
			return nil, err
		}
		slugs = append(slugs, slug)
	}
	return slugs, nil
}
