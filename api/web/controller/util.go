package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getUrlParam(r *http.Request, name string) string {
	vars := mux.Vars(r)
	return vars[name]
}
