package handlers

import (
	"net/http"

	"github.com/Omelman/trucking-api/src/service"
)

type ShipmentHandler struct {
	service *service.Service
}

func NewShipmentHandler(s *service.Service) *ShipmentHandler {
	return &ShipmentHandler{
		service: s,
	}
}

func (h *ShipmentHandler) CreateShipment(w http.ResponseWriter, r *http.Request) {
	// logic
}

func (h *ShipmentHandler) UpdateShipment(w http.ResponseWriter, r *http.Request) {
	// logic
}

func (h *ShipmentHandler) GetShipment(w http.ResponseWriter, r *http.Request) {
	// logic
}

func (h *ShipmentHandler) GetAllCustomerShipment(w http.ResponseWriter, r *http.Request) {
	// logic
}
