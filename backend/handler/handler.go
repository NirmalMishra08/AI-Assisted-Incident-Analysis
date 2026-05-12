package handler

import "backend/handler/auth"


type Handler struct {
	Auth *auth.Handler
}

func New() *Handler{
	return &Handler{
		Auth: auth.New(),
	}
}