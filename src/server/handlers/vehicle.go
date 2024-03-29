package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Omelman/trucking-api/src/context"
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

	userID := context.GetUserID(r.Context())
	req.OwnerID = userID

	err = h.service.CreateVehicle(r.Context(), req)
	if err != nil {
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendEmptyResponse(w, http.StatusCreated)
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
	resp, err := h.service.GetAllVehicles(r.Context())
	if err != nil {
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendResponse(w, http.StatusOK, resp)
}

func (h *VehicleHandler) GetAllOwnerVehicle(w http.ResponseWriter, r *http.Request) {
	userID := context.GetUserID(r.Context())

	resp, err := h.service.GetUserVehicles(r.Context(), userID)
	if err != nil {
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendResponse(w, http.StatusOK, resp)
}

func (h *VehicleHandler) DeleteUserVehicle(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["vehicle_id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		SendEmptyResponse(w, http.StatusBadRequest)

		return
	}

	userID := context.GetUserID(r.Context())

	err = h.service.DeleteUserVehicles(r.Context(), id, userID)
	if err != nil {
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendEmptyResponse(w, http.StatusOK)
}

func (h *VehicleHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.GetStatistics(r.Context())
	if err != nil {
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendResponse(w, http.StatusOK, resp)
}

func (h *VehicleHandler) CreateShipment(w http.ResponseWriter, r *http.Request) {
	req := &models.Shipment{}

	err := UnmarshalRequest(r, req)
	if err != nil {
		SendEmptyResponse(w, http.StatusBadRequest)

		return
	}

	resp, err := h.service.CreateShipment(r.Context(), req)
	if err != nil {
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendResponse(w, http.StatusOK, resp)
}

func (h *VehicleHandler) UpdateShipment(w http.ResponseWriter, r *http.Request) {
	req := &models.Shipment{}

	err := UnmarshalRequest(r, req)
	if err != nil {
		SendEmptyResponse(w, http.StatusBadRequest)

		return
	}

	err = h.service.UpdateShipment(r.Context(), req)
	if err != nil {
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendEmptyResponse(w, http.StatusOK)
}
