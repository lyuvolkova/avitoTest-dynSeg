package server

import (
	"net/http"
)

// swagger:parameters getUserSegments
type GetUserSegmentsRequest struct {
	// in: path
	// example: 89
	UserID int64 `json:"userID"`
}

// swagger:response getUserSegmentsResponse
type GetUserSegmentsResponse struct {
	// example: ["avito_discount_70", "avito_discount_50"]
	Slugs []string `json:"slugs"`
}

// swagger:route GET /users/{userID}/segments user-segments getUserSegments
//
// Возвращает сегменты, в которых состоит пользователь
//
//		Responses:
//		  200: getUserSegmentsResponse
//	      500: errorResponse
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
