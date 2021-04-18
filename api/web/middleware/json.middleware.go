package middleware

import (
	"errors"
	"net/http"
	"strings"

	"Sharykhin/rent-car/api/web/response"
	"Sharykhin/rent-car/domain"
)

func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if (r.Method == "POST" || r.Method == "PUT") &&
			!strings.Contains(r.Header.Get("content-type"), "application/json") {

			response.Fail(
				w,
				domain.NewError(
					errors.New("Content-type must be application/json."),
					domain.ValidationErrorCode,
					"Content-type must be application/json.",
				))
			return
		}

		next.ServeHTTP(w, r)
	})
}
