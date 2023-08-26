package server

type handler struct {
	app app
}

func newHandler(app app) *handler {

	return &handler{app: app}

}
