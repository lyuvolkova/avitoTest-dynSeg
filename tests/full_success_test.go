package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lyuvolkova/avitoTest-dynSeg/tests/client/client"
	"github.com/lyuvolkova/avitoTest-dynSeg/tests/client/client/segments"
	"github.com/lyuvolkova/avitoTest-dynSeg/tests/client/client/user_segments"
	"github.com/lyuvolkova/avitoTest-dynSeg/tests/client/models"
)

func TestFullSuccess(t *testing.T) {
	var err error

	r := require.New(t)
	c := client.Default
	ctx := context.Background()

	slugs := []string{"AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS", "AVITO_DISCOUNT_30", "AVITO_DISCOUNT_50"}

	for _, v := range slugs {
		_, err = c.Segments.DeleteSegment(&segments.DeleteSegmentParams{Context: ctx, Slug: v})
		r.NoError(err)
	}

	for _, v := range slugs {
		_, err = c.Segments.CreateSegment(&segments.CreateSegmentParams{
			Context: ctx,
			Body: &models.CreateSegmentRequest{
				Slug: v,
			},
		})
		r.NoError(err)
	}

	usersSegments := map[int64][]string{
		1000: {"AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS", "AVITO_DISCOUNT_30"},
		1002: {"AVITO_VOICE_MESSAGES", "AVITO_DISCOUNT_30"},
		1004: nil,
	}

	for userID, slugs := range usersSegments {
		if len(slugs) == 0 {
			continue
		}

		_, err = c.UserSegments.UpdateUserSegments(&user_segments.UpdateUserSegmentsParams{
			Context: ctx,
			Body: &models.UpdateUserSegmentsRequest{
				AddSlug: slugs,
			}, UserID: userID})
		r.NoError(err)
	}

	for userID, slugs := range usersSegments {
		resp, err := c.UserSegments.GetUserSegments(&user_segments.GetUserSegmentsParams{
			Context: ctx,
			UserID:  userID})
		r.NoError(err)
		r.EqualValues(slugs, resp.Payload.Slugs)
	}

	_, err = c.UserSegments.UpdateUserSegments(&user_segments.UpdateUserSegmentsParams{
		Context: ctx,
		Body: &models.UpdateUserSegmentsRequest{
			AddSlug:    []string{"AVITO_DISCOUNT_50"},
			DeleteSlug: []string{"AVITO_PERFORMANCE_VAS", "AVITO_DISCOUNT_30"},
		}, UserID: 1000})
	r.NoError(err)

	expSlugs := []string{"AVITO_VOICE_MESSAGES", "AVITO_DISCOUNT_50"}
	resp, err := c.UserSegments.GetUserSegments(&user_segments.GetUserSegmentsParams{
		Context: ctx,
		UserID:  1000})
	r.NoError(err)
	r.EqualValues(expSlugs, resp.Payload.Slugs)

	for _, v := range slugs {
		_, err = c.Segments.DeleteSegment(&segments.DeleteSegmentParams{Context: ctx, Slug: v})
		r.NoError(err)
	}
}
