package handlers

import "dogovorsBackEnd/internal/use/controllers"

type Handler interface {
}

type handler struct {
	controller controllers.Controller
}

func New(controller controllers.Controller) Handler {
	return &handler{controller: controller}
}
