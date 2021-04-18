package controller

import (
	"encoding/json"
	"net/http"

	"Sharykhin/rent-car/api/web/response"
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/services"
)

type (
	ConsumerController struct {
		consumerService *services.ConsumerService
		logger          domain.LoggerInterface
	}

	CreateConsumerPayload struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

func NewConsumerController(consumerService *services.ConsumerService, logger domain.LoggerInterface) *ConsumerController {
	ctrl := ConsumerController{
		consumerService: consumerService,
		logger:          logger,
	}

	return &ctrl
}

func (c *ConsumerController) CreateConsumer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload CreateConsumerPayload
	err := decoder.Decode(&payload)

	if err != nil {
		c.logger.Error(err.Error())
		response.Fail(w, err)
		return
	}

	consumer, err := c.consumerService.CreateNewConsumer(r.Context(), payload.FirstName, payload.LastName, payload.Email)
	if err != nil {
		c.logger.Error(err.Error())
		response.Fail(w, err)
		return
	}

	response.Created(w, consumer, nil)
}
