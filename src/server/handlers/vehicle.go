package handlers

import (
	"net/http"

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
	// logic
}

func (h *VehicleHandler) GetVehicle(w http.ResponseWriter, r *http.Request) {
	// logic
}

func (h *VehicleHandler) GetAllOwnerVehicle(w http.ResponseWriter, r *http.Request) {
	// logic
}
