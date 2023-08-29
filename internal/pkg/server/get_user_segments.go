package server

import (
	"net/http"
)

type GetUserSegmentsRequest struct {
	UserID int64 `json:"-"`
}

type GetUserSegmentsResponse struct {
	Slugs []string `json:"slugs"`
}

func (h *handler) getUserSegments(w http.ResponseWriter, r *http.Request) {
	request := GetUserSegmentsRequest{}

	var err error

	request.UserID, err = PathVarInt64(r, "userID")
	if err != nil {
		sendResponse(nil, ErrBadRequest, w)

		return
	}

	response, err := h.app.GetUserSegments(r.Context(), &request)
	sendResponse(response, err, w)
}
