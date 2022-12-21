package controller

import (
	"github.com/jsovalles/rest-ws/internal/utils"
	"net/http"
)

type HomeController interface {
	GetHome(w http.ResponseWriter, r *http.Request)
}

type homeController struct {
	env utils.Env
}

func (h *homeController) GetHome(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewHomeController(env utils.Env) HomeController {
	return &homeController{env: env}
}
