package middleware

import "net/http"

type (
	Handler func(http.Handler) http.Handler
)
