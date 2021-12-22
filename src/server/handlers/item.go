package handlers

import (
	"net/http"

	"github.com/Omelman/trucking-api/src/models"
	"github.com/Omelman/trucking-api/src/service"
)

type ItemHandler struct {
	service *service.Service
}

func NewItemHandler(s *service.Service) *ItemHandler {
	return &ItemHandler{
		service: s,
	}
}

func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	req := &models.Item{}

	err := UnmarshalRequest(r, req)
	if err != nil {
		SendEmptyResponse(w, http.StatusBadRequest)

		return
	}

	err = h.service.CreateItem(r.Context(), req)
	if err != nil {
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendEmptyResponse(w, http.StatusCreated)
}

func (h *ItemHandler) UpdateShipment(w http.ResponseWriter, r *http.Request) {
	// logic
}

func (h *ItemHandler) GetShipment(w http.ResponseWriter, r *http.Request) {
	// logic
}

func (h *ItemHandler) GetAllCustomerShipment(w http.ResponseWriter, r *http.Request) {
	// logic
}
