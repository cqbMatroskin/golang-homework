package auth

import (
	"fmt"
	"net/http"
	"validation-api/configs"
	"validation-api/pkg/res"
)

type AuthHandler struct{
	*configs.AuthConfig
}

type AuthHandlerDeps struct{
	*configs.AuthConfig
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		AuthConfig: deps.AuthConfig,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request)  {
		fmt.Println("Login")
		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
	}
}