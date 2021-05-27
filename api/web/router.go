package web

import (
	"net/http"

	"github.com/gorilla/mux"

	"Sharykhin/rent-car/api/web/middleware"
	"Sharykhin/rent-car/di"
)

func router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/_healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}).Methods("GET")

	sr := r.PathPrefix("/v1").Subrouter()
	sr.Use(middleware.LoggingMiddleware(di.Container.Logger), middleware.JsonMiddleware)

	carController := di.Container.CarController
	consumerController := di.Container.ConsumerController
	requisitionController := di.Container.RequisitionController

	sr.HandleFunc("/cars", carController.CreateCar).Methods("POST")
	sr.HandleFunc("/cars/{id}", carController.GetCarByID).Methods("GET")
	sr.HandleFunc("/consumers", consumerController.CreateConsumer).Methods("POST")
	sr.HandleFunc("/requisitions", requisitionController.CreateRequisition).Methods("POST")

	return r
}
