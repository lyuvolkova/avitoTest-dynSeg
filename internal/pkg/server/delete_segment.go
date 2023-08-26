package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type DeleteSegmentRequest struct {
	Slug string `json:"-"`
}

type DeleteSegmentResponse struct {
	Ok bool `json:"ok"`
}

func (h *handler) deleteSegment(w http.ResponseWriter, r *http.Request) {
	request := DeleteSegmentRequest{
		Slug: mux.Vars(r)["slug"],
	}

	response, err := h.app.DeleteSegment(r.Context(), &request)
	sendResponse(response, err, w)
}
