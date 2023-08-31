package server

import (
	"encoding/json"
	"net/http"
)

// swagger:parameters createSegment
type swaggerCreateSegmentRequest struct {
	// in: body
	Body CreateSegmentRequest `json:"body"`
}

type CreateSegmentRequest struct {
	Slug string `json:"slug"`
}

// swagger:response createSegmentResponse
type CreateSegmentResponse struct {
	Ok bool `json:"ok"`
}

// swagger:route POST /segments segments createSegment
//
// Создает новый сегмент
//
//		Responses:
//		  200: createSegmentResponse
//	      500: errorResponse
func (h *handler) createSegment(w http.ResponseWriter, r *http.Request) {
	request := CreateSegmentRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendResponse(nil, ErrBadRequest, w)

		return
	}

	response, err := h.app.CreateSegment(r.Context(), &request)
	sendResponse(response, err, w)
}
