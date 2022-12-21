package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jsovalles/rest-ws/internal/middleware"
	"github.com/jsovalles/rest-ws/internal/models"
	"github.com/jsovalles/rest-ws/internal/repository"
	"github.com/jsovalles/rest-ws/internal/utils"
	"github.com/jsovalles/rest-ws/internal/websocket"
	"github.com/segmentio/ksuid"
	"net/http"
	"strconv"
	"time"
)

type PostController interface {
	CreatePost(w http.ResponseWriter, r *http.Request)
	GetPostById(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
	ListPosts(w http.ResponseWriter, r *http.Request)
}

type postController struct {
	env        utils.Env
	repository repository.PostRepository
	middleware middleware.Auth
	util       utils.Util
	hub        websocket.Hub
}

func NewPostController(env utils.Env, repository repository.PostRepository, middleware middleware.Auth, util utils.Util, hub websocket.Hub) PostController {
	return &postController{env: env, repository: repository, middleware: middleware, util: util, hub: hub}
}

func (p *postController) CreatePost(w http.ResponseWriter, r *http.Request) {
	token, err := p.middleware.GetAuthToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		var postRequest models.CreatePostRequest
		err := json.NewDecoder(r.Body).Decode(&postRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		post := models.Post{
			Id:          id.String(),
			PostContent: postRequest.PostContent,
			CreatedAt:   time.Now(),
			UserId:      claims.UserId,
		}

		err = p.repository.CreatePost(post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var postMessage = models.WebsocketMessage{Type: "PostCreated", Payload: post}

		p.hub.Broadcast(postMessage, nil)

		p.util.WriteJSON(w, http.StatusOK, models.PostResponse{
			Id:          post.Id,
			PostContent: post.PostContent,
		}, nil)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *postController) GetPostById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	post, err := p.repository.GetPostById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.util.WriteJSON(w, http.StatusOK, models.PostResponse{
		Id:          post.Id,
		PostContent: post.PostContent,
	}, nil)
}

func (p *postController) UpdatePost(w http.ResponseWriter, r *http.Request) {
	token, err := p.middleware.GetAuthToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if _, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		params := mux.Vars(r)
		id := params["id"]
		var postRequest models.CreatePostRequest
		err := json.NewDecoder(r.Body).Decode(&postRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		post, err := p.repository.GetPostById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		post.PostContent = postRequest.PostContent

		err = p.repository.UpdatePost(post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		p.util.WriteJSON(w, http.StatusOK, models.PostResponse{
			Id:          post.Id,
			PostContent: post.PostContent,
		}, nil)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *postController) DeletePost(w http.ResponseWriter, r *http.Request) {
	token, err := p.middleware.GetAuthToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		params := mux.Vars(r)
		id := params["id"]

		_, err := p.repository.GetPostById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = p.repository.DeletePostByIdAndUserId(id, claims.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		p.util.WriteJSON(w, http.StatusOK, "deleted", nil)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *postController) ListPosts(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	var err error
	var page = uint64(0)
	if pageStr != "" {
		page, err = strconv.ParseUint(pageStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	posts, err := p.repository.ListPosts(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.util.WriteJSON(w, http.StatusOK, posts, nil)
}
