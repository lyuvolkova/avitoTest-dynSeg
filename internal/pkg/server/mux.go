package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewMux(app app) http.Handler {
	h := newHandler(app)
	router := mux.NewRouter()

	router.Path("/segments").Methods(http.MethodPost).HandlerFunc(h.createSegment)
	router.Path("/segments/{slug:[a-zA-Z0-9-_]+}").Methods(http.MethodDelete).HandlerFunc(h.deleteSegment)
	router.Path("/users/{userID:[0-9]+}/segments").Methods(http.MethodPatch).HandlerFunc(h.updateUserSegments)
	router.Path("/users/{userID:[0-9]+}/segments").Methods(http.MethodGet).HandlerFunc(h.getUserSegments)

	return router
}
