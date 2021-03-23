package web

import (
	"Sharykhin/rent-car/api/web/controller"
	"Sharykhin/rent-car/domain/car/services"
	"Sharykhin/rent-car/infrastructure/postgres"
	"Sharykhin/rent-car/infrastructure/postgres/repositories"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/_healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}).Methods("GET")

	sr := r.PathPrefix("/v1").Subrouter()
	db, err := postgres.Connect(os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	carController := controller.NewCarController(services.NewCarService(repositories.NewCarRepository(db)))
	sr.HandleFunc("/cars", carController.CreateCar).Methods("POST")

	return r
}
