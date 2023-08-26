package server

import (
	"encoding/json"
	"net/http"
)

type CreateSegmentRequest struct {
	Slug string `json:"slug"`
}

type CreateSegmentResponse struct {
	Ok bool `json:"ok"`
}

func (h *handler) createSegment(w http.ResponseWriter, r *http.Request) {
	request := CreateSegmentRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendResponse(nil, ErrBadRequest, w)

		return
	}

	response, err := h.app.CreateSegment(r.Context(), &request)
	sendResponse(response, err, w)
}
