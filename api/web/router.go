package web

import (
	"Sharykhin/rent-car/api/web/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/_healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}).Methods("GET")

	sr := r.PathPrefix("/v1").Subrouter()

	carController := controller.CarController{}
	sr.HandleFunc("/cars", carController.CreateCar).Methods("POST")

	return r
}
