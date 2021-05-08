package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Method: %s, URI: %s", r.Method, r.RequestURI)

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Errorf("Error reading request body: %v", err.Error())

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Debugf("Request body: %+v", string(buf))

		reader := ioutil.NopCloser(bytes.NewBuffer(buf))
		r.Body = reader

		next.ServeHTTP(w, r)
	})
}
