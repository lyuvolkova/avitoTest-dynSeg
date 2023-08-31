package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// swagger:parameters deleteSegment
type DeleteSegmentRequest struct {
	// in: path
	Slug string `json:"slug"`
}

// swagger:response deleteSegmentResponse
type DeleteSegmentResponse struct {
	Ok bool `json:"ok"`
}

// swagger:route DELETE /segments/{slug} segments deleteSegment
//
// Удаляет сегмент
//
//		Responses:
//		  200: deleteSegmentResponse
//	      500: errorResponse
func (h *handler) deleteSegment(w http.ResponseWriter, r *http.Request) {
	request := DeleteSegmentRequest{
		Slug: mux.Vars(r)["slug"],
	}

	response, err := h.app.DeleteSegment(r.Context(), &request)
	sendResponse(response, err, w)
}
