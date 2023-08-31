package server

import (
	"encoding/json"
	"net/http"
)

// swagger:parameters updateUserSegments
type swaggerUpdateUserSegmentsRequest struct {
	// in: body
	Body UpdateUserSegmentsRequest `json:"body"`

	// in: path
	UserID int64 `json:"userID"`
}

type UpdateUserSegmentsRequest struct {
	// добавляет пользователя в эти сегменты
	AddSlug []string `json:"add_slug"`

	// удаляет пользователя из этих сегментов
	DeleteSlug []string `json:"delete_slug"`

	UserID int64 `json:"-"`
}

// swagger:response updateUserSegmentsResponse
type UpdateUserSegmentsResponse struct {
	Ok bool `json:"ok"`
}

// swagger:route PATCH /users/{userID}/segments user-segments updateUserSegments
//
// Обновляет сегменты у пользователя
//
//		Responses:
//		  200: updateUserSegmentsResponse
//	      500: errorResponse
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
