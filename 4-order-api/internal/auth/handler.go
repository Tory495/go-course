package auth

import (
	"go/courses/configs"
	"go/courses/pkg/generator"
	"go/courses/pkg/jwt"
	"go/courses/pkg/request"
	"go/courses/pkg/response"
	"net/http"
)

type AuthHandler struct {
	Config      *configs.Config
	AuthService *AuthService
}

type AuthHandlerDeps struct {
	Config      *configs.Config
	AuthService *AuthService
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}

	router.HandleFunc("POST /auth", handler.auth())
	router.HandleFunc("POST /confirm", handler.confirm())
}

func (h *AuthHandler) auth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[AuthRequest](&w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sessionId, err := generator.GenerateSessionId()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		auth := &Auth{
			Phone:     body.Phone,
			SessionId: sessionId,
			Code:      0,
		}

		_, err = h.AuthService.AuthRepository.StoreAuth(auth)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = h.AuthService.SendSms(body.Phone)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.SendJsonResponse(w, AuthResponse{
			SessionId: sessionId,
		}, http.StatusOK)
	}
}

func (h *AuthHandler) confirm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[ConfirmRequest](&w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		auth, err := h.AuthService.AuthRepository.GetBySessionId(body.SessionId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if auth.Code != body.Code {
			http.Error(w, "Codes mismatch", http.StatusBadRequest)
			return
		}

		token, err := jwt.NewJWT(h.Config.Auth.Secret).Create(auth.Phone)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.SendJsonResponse(w, token, http.StatusOK)
	}
}
