package middleware

import (
	"net/http"
	"strings"

	"Sharykhin/rent-car/api/web/response"
	"Sharykhin/rent-car/domain"
)

func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if (r.Method == "POST" || r.Method == "PUT") &&
			!strings.Contains(r.Header.Get("content-type"), "application/json") {
			response.Fail(w, "Content-type must be application/json.", domain.ValidationErrorCode)
			return
		}

		next.ServeHTTP(w, r)
	})
}
