package middleware

import "golink/internal/infra/env"

type Middleware struct {
	env *env.ENV
}

func NewMiddleware(env *env.ENV) *Middleware {
	return &Middleware{env: env}
}
