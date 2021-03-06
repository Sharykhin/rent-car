package middleware

import (
	"errors"
	"net/http"
	"strings"

	"Sharykhin/rent-car/api/web/response"
	"Sharykhin/rent-car/domain"
)

// JsonMiddleware checks that incoming request has content-type header that equals application/json
func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if (r.Method == "POST" || r.Method == "PUT") &&
			!strings.Contains(r.Header.Get("content-type"), "application/json") {

			response.Fail(
				w,
				domain.NewError(
					errors.New("content-type must be application/json"),
					"[api][web][middleware][JsonMiddleware]",
					domain.ValidationErrorCode,
				))
			return
		}

		next.ServeHTTP(w, r)
	})
}
