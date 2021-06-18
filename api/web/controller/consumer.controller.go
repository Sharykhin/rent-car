package controller

import (
	"encoding/json"
	"net/http"

	"Sharykhin/rent-car/api/web/response"
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/service"
)

type (
	// ConsumerController handle rest api endpoint around consumer
	ConsumerController struct {
		consumerService *service.ConsumerService
	}

	// CreateConsumerPayload is a payload for creating a new consumer
	CreateConsumerPayload struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

// NewConsumerController return a new instance of consumer controller
func NewConsumerController(consumerService *service.ConsumerService) *ConsumerController {
	ctrl := ConsumerController{
		consumerService: consumerService,
	}

	return &ctrl
}

// CreateConsumer handles endpoint for creating a new consumer
func (c *ConsumerController) CreateConsumer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload CreateConsumerPayload
	err := decoder.Decode(&payload)

	if err != nil {
		response.Fail(w, domain.WrapErrorWithStack(err, "[api][web][controller][ConsumerController][CreateConsumer]"))
		return
	}

	consumer, err := c.consumerService.CreateNewConsumer(r.Context(), payload.FirstName, payload.LastName, payload.Email)
	if err != nil {
		response.Fail(w, domain.WrapErrorWithStack(err, "[api][web][controller][ConsumerController][CreateConsumer]"))
		return
	}

	response.Created(w, consumer, nil)
}
