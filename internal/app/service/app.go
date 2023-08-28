package service

type app struct {
	sp segmentsRepo
}

func New(sp segmentsRepo) *app {
	return &app{sp: sp}
}
