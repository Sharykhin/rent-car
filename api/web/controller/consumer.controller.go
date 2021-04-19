package controller

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"Sharykhin/rent-car/api/web/response"
	"Sharykhin/rent-car/domain/consumer/services"
)

type (
	ConsumerController struct {
		consumerService *services.ConsumerService
	}

	CreateConsumerPayload struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

func NewConsumerController(consumerService *services.ConsumerService) *ConsumerController {
	ctrl := ConsumerController{
		consumerService: consumerService,
	}

	return &ctrl
}

func (c *ConsumerController) CreateConsumer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload CreateConsumerPayload
	err := decoder.Decode(&payload)

	if err != nil {
		log.Error(err)
		response.Fail(w, err)
		return
	}

	consumer, err := c.consumerService.CreateNewConsumer(r.Context(), payload.FirstName, payload.LastName, payload.Email)
	if err != nil {
		log.Error(err)
		response.Fail(w, err)
		return
	}

	response.Created(w, consumer, nil)
}
