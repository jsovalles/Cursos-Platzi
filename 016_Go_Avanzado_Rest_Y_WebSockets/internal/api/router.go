package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jsovalles/rest-ws/internal/controller"
	"github.com/jsovalles/rest-ws/internal/middleware"
	"github.com/jsovalles/rest-ws/internal/utils"
	"github.com/jsovalles/rest-ws/internal/websocket"
	"github.com/rs/cors"
	"go.uber.org/fx"
	"net/http"
)

type Api interface {
	SetupRoutes()
}

type api struct {
	homeController controller.HomeController
	userController controller.UserController
	postController controller.PostController
	middleware     middleware.Auth
	env            utils.Env
	hub            websocket.Hub
}

func (a *api) SetupRoutes() {
	r := mux.NewRouter()
	apiRoute := r.PathPrefix("/api").Subrouter().StrictSlash(false)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
	})

	apiRoute.Use(a.middleware.CheckAuthMiddleware())
	apiRoute.HandleFunc("/hello", a.homeController.GetHome).Methods(http.MethodGet)
	apiRoute.HandleFunc("/users/signup", a.userController.SignUp).Methods(http.MethodPost)
	apiRoute.HandleFunc("/users/login", a.userController.Login).Methods(http.MethodPost)
	apiRoute.HandleFunc("/users/me", a.userController.MeHandler).Methods(http.MethodGet)
	apiRoute.HandleFunc("/posts", a.postController.CreatePost).Methods(http.MethodPost)
	apiRoute.HandleFunc("/posts/{id}", a.postController.GetPostById).Methods(http.MethodGet)
	apiRoute.HandleFunc("/posts/{id}", a.postController.UpdatePost).Methods(http.MethodPut)
	apiRoute.HandleFunc("/posts/{id}", a.postController.DeletePost).Methods(http.MethodDelete)
	apiRoute.HandleFunc("/posts", a.postController.ListPosts).Methods(http.MethodGet)
	apiRoute.HandleFunc("/ws", a.hub.HandlerWebSocket)
	if a.env.Environment == "Local" {
		fmt.Println("swagger")
		//apiRoute.PathPrefix("/").Handler(httpSwagger.WrapHandler)
	}

	if err := http.ListenAndServe(":"+a.env.Port, c.Handler(r)); err != nil {
		fmt.Errorf(fmt.Sprintf("%s%s", "Failed to listen and serve on port ", a.env.Port) + err.Error())
		panic(err)
	}
}

func NewRoutes(homeController controller.HomeController, userController controller.UserController, postController controller.PostController, middleware middleware.Auth, env utils.Env, hub websocket.Hub) Api {
	return &api{homeController: homeController, userController: userController, postController: postController, middleware: middleware, env: env, hub: hub}
}

var Module = fx.Provide(NewRoutes)
