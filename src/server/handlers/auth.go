package handlers

import (
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Omelman/trucking-api/src/models"
	"github.com/Omelman/trucking-api/src/service"
)

type AuthHandler struct {
	service *service.Service
}

func NewAuthHandler(s *service.Service) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

// swagger:operation POST /login auth login
//   create a session and obtain tokens pair
// ---
// parameters:
// - name: login_request
//   in: body
//   required: true
//   schema:
//     $ref: '#/definitions/LoginRequest'
// responses:
//   '200':
//     description: Fetched
//     schema:
//       "$ref": "#/definitions/LoginResponse"
//   '400':
//     description: Bad Request
//     schema:
//       "$ref": "#/definitions/ValidationErr"
//   '500':
//     description: Internal Server Error
//     schema:
//       "$ref": "#/definitions/CommonError"
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := &models.LoginRequest{}

	err := UnmarshalRequest(r, req)
	if err != nil {
		SendEmptyResponse(w, http.StatusBadRequest)

		return
	}

	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	res, err := h.service.Login(r.Context(), *req)
	if err != nil {
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendResponse(w, http.StatusOK, res)
}

// swagger:operation DELETE /logout auth logout
//   deactivate user session, move access token to the black list
// ---
// responses:
//   '204':
//     description: Successfully logged out
//   '400':
//     description: Bad Request
//     schema:
//       "$ref": "#/definitions/ValidationErr"
//   '500':
//     description: Internal Server Error
//     schema:
//       "$ref": "#/definitions/CommonError"
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	jwtAccessToken := r.Header.Get(AccessTokenHeader)

	err := h.service.Logout(r.Context(), jwtAccessToken)
	if err != nil {
		log.Error("logout error ", err)
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendResponse(w, http.StatusOK, nil)
}

// swagger:operation POST /token auth token
//   refresh access token if previous tokens pair was valid
// ---
// parameters:
// - name: TokenPair
//   in: body
//   required: true
//   schema:
//     $ref: '#/definitions/TokenPair'
// responses:
//   '201':
//     description: Created
//     schema:
//       "$ref": "#/definitions/TokenPair"
//   '400':
//     description: Bad Request
//     schema:
//       "$ref": "#/definitions/ValidationErr"
//   '500':
//     description: Internal Server Error
//     schema:
//       "$ref": "#/definitions/CommonError"
func (h *AuthHandler) TokenRefresh(w http.ResponseWriter, r *http.Request) {
	req := &models.TokenPair{}

	err := UnmarshalRequest(r, req)
	if err != nil {
		log.Error("token refresh error ", err)
		SendEmptyResponse(w, http.StatusBadRequest)

		return
	}

	res, err := h.service.RefreshToken(r.Context(), req)
	if err != nil {
		log.Error("token refresh error ", err)
		SendEmptyResponse(w, http.StatusInternalServerError)

		return
	}

	SendResponse(w, http.StatusCreated, res)
}

// swagger:operation POST /users/register auth create-user
// ---
// parameters:
// - name: Request
//   in: body
//   required: true
//   schema:
//     $ref: '#/definitions/UserRegistrationRequest'
// responses:
//   '201':
//     description: Created
//   '400':
//     description: Bad Request
//     schema:
//       "$ref": "#/definitions/ValidationErr"
//   '500':
//     description: Internal Server Error
//     schema:
//       "$ref": "#/definitions/CommonError"
func (h *AuthHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		usrReq models.UserRegistrationRequest
		ctx    = r.Context()
	)

	err := UnmarshalRequest(r, &usrReq)
	if err != nil {
		SendEmptyResponse(w, http.StatusBadRequest)
		return
	}

	usrReq.Email = strings.ToLower(strings.TrimSpace(usrReq.Email))

	user := &models.User{
		FirstName: usrReq.FirstName,
		LastName:  usrReq.LastName,
		Password:  usrReq.Password,
		Email:     usrReq.Email,
		RoleID:    usrReq.RoleID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err = h.service.CreateUser(ctx, user)
	if err != nil {
		log.WithContext(ctx).Error(ctx, err, "create user")
		SendHTTPError(w, err)

		return
	}

	SendEmptyResponse(w, http.StatusCreated)
}
