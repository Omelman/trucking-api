package handlers

import (
	"net/http"

	"github.com/Omelman/trucking-api/src/context"
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

	userID := context.GetUserID(r.Context())
	req.UserID = userID

	err = h.service.CreateItem(r.Context(), req)
	if err != nil {
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendEmptyResponse(w, http.StatusCreated)
}

func (h *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	req := &models.Item{}

	err := UnmarshalRequest(r, req)
	if err != nil {
		SendEmptyResponse(w, http.StatusBadRequest)

		return
	}

	userID := context.GetUserID(r.Context())
	req.UserID = userID

	err = h.service.UpdateItem(r.Context(), req)
	if err != nil {
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendEmptyResponse(w, http.StatusCreated)
}

func (h *ItemHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.GetAllItems(r.Context())
	if err != nil {
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendResponse(w, http.StatusOK, resp)
}
