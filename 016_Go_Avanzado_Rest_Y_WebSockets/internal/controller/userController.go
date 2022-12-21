package controller

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/jsovalles/rest-ws/internal/middleware"
	"github.com/jsovalles/rest-ws/internal/models"
	"github.com/jsovalles/rest-ws/internal/repository"
	"github.com/jsovalles/rest-ws/internal/utils"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type UserController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	MeHandler(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	env        utils.Env
	repository repository.UserRepository
	middleware middleware.Auth
}

func (u *userController) SignUp(w http.ResponseWriter, r *http.Request) {
	var request models.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//This logic should be on service layer
	id, err := ksuid.NewRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user = models.User{
		Id:       id.String(),
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	err = u.repository.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.SignUpResponse{
		Id:    user.Id,
		Email: user.Email,
	})
}

func (u *userController) Login(w http.ResponseWriter, r *http.Request) {
	var request models.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := u.repository.GetUserByEmail(request.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	claims := models.AppClaims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(u.env.JwtSecret))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.LoginResponse{
		Token: tokenString,
	})
}

func (u *userController) MeHandler(w http.ResponseWriter, r *http.Request) {
	token, err := u.middleware.GetAuthToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		user, err := u.repository.GetUserById(claims.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewUserController(env utils.Env, repository repository.UserRepository, middleware middleware.Auth) UserController {
	return &userController{env: env, repository: repository, middleware: middleware}
}
