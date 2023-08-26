package server

import (
	"encoding/json"
	"net/http"
)

type UpdateUserSegmentsRequest struct {
	AddSlug    []string `json:"add_slug"`
	DeleteSlug []string `json:"delete_slug"`
	UserID     int64    `json:"-"`
}

type UpdateUserSegmentsResponse struct {
	Ok bool `json:"ok"`
}

func (h *handler) updateUserSegments(w http.ResponseWriter, r *http.Request) {
	request := UpdateUserSegmentsRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendResponse(nil, ErrBadRequest, w)

		return
	}

	var err error

	request.UserID, err = PathVarInt64(r, "userID")
	if err != nil {
		sendResponse(nil, ErrBadRequest, w)

		return
	}

	response, err := h.app.UpdateUserSegments(r.Context(), &request)
	sendResponse(response, err, w)
}
