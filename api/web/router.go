package web

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"Sharykhin/rent-car/api/web/controller"
	"Sharykhin/rent-car/api/web/middleware"
	carServices "Sharykhin/rent-car/domain/car/services"
	consumerServices "Sharykhin/rent-car/domain/consumer/services"
	"Sharykhin/rent-car/infrastructure/postgres"
	"Sharykhin/rent-car/infrastructure/postgres/repositories"
)

func router() http.Handler {
	db, err := postgres.Connect(os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/_healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}).Methods("GET")

	sr := r.PathPrefix("/v1").Subrouter()
	sr.Use(middleware.LoggingMiddleware, middleware.JsonMiddleware)

	carController := controller.NewCarController(carServices.NewCarService(repositories.NewCarRepository(db)))
	consumerController := controller.NewConsumerController(consumerServices.NewConsumerService(repositories.NewConsumerRepository(db)))

	sr.HandleFunc("/cars", carController.CreateCar).Methods("POST")
	sr.HandleFunc("/cars/{id}", carController.GetCarByID).Methods("GET")
	sr.HandleFunc("/consumers", consumerController.CreateConsumer).Methods("POST")

	return r
}
