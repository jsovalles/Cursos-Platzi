package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/jsovalles/rest-ws/internal/models"
	"github.com/jsovalles/rest-ws/internal/utils"
	"go.uber.org/fx"
	"net/http"
	"strings"
)

type Auth interface {
	CheckAuthMiddleware() func(h http.Handler) http.Handler
	GetAuthToken(r *http.Request) (token *jwt.Token, err error)
}

type auth struct {
	env utils.Env
}

var (
	NO_AUTH_NEEDED = []string{
		"login",
		"signup",
	}
)

func shouldCheckToken(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

func (a *auth) CheckAuthMiddleware() func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})

	}
}

func (a *auth) GetAuthToken(r *http.Request) (token *jwt.Token, err error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
	token, err = jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.env.JwtSecret), nil
	})

	if err != nil {
		return
	}

	return
}

func NewAuth(env utils.Env) Auth {
	return &auth{env: env}
}

var Module = fx.Provide(NewAuth)
