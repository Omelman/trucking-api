package handlers

import (
	"net/http"

	"github.com/Omelman/trucking-api/src/models"
	"github.com/Omelman/trucking-api/src/service"
)

type VehicleHandler struct {
	service *service.Service
}

func NewVehicleHandler(s *service.Service) *VehicleHandler {
	return &VehicleHandler{
		service: s,
	}
}

func (h *VehicleHandler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	req := &models.Vehicle{}

	err := UnmarshalRequest(r, req)
	if err != nil {
		SendEmptyResponse(w, http.StatusBadRequest)

		return
	}

	err = h.service.CreateVehicle(r.Context(), req)
	if err != nil {
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendEmptyResponse(w, http.StatusOK)
}

func (h *VehicleHandler) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	req := &models.Vehicle{}

	err := UnmarshalRequest(r, req)
	if err != nil {
		SendEmptyResponse(w, http.StatusBadRequest)

		return
	}

	err = h.service.UpdateVehicle(r.Context(), req)
	if err != nil {
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendEmptyResponse(w, http.StatusOK)
}

func (h *VehicleHandler) GetVehicle(w http.ResponseWriter, r *http.Request) {
	// logic
}

func (h *VehicleHandler) GetAllOwnerVehicle(w http.ResponseWriter, r *http.Request) {
	// logic
}
