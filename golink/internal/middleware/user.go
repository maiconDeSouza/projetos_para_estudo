package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golink/internal/infra/env"
	"golink/internal/infra/errapp"
	"golink/internal/service"
	"net/http"
)

func (m *Middleware) AuthUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		cookie, err := r.Cookie("session_token_user")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				errResp := errapp.NewUnauthorizedError("Token de sessão não fornecido", err)
				w.WriteHeader(errResp.Code)
				json.NewEncoder(w).Encode(errResp)
				return
			}
			errResp := errapp.NewInternalServerError("Erro do servidor", err)
			w.WriteHeader(errResp.Code)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		tokenString := cookie.Value

		sub, isAdmin, err := service.ValidarToken(tokenString, m.env.JWT_SECRET_KEY)
		fmt.Println(sub)
		if err != nil {
			errResp := errapp.NewUnauthorizedError(err.Error(), err)
			w.WriteHeader(errResp.Code)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		if isAdmin {
			errResp := errapp.NewUnauthorizedError("Não autorizado", nil)
			w.WriteHeader(errResp.Code)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		ctx := context.WithValue(r.Context(), env.UserSubKey, sub)
		ctx = context.WithValue(ctx, env.UserAdminKey, isAdmin)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
