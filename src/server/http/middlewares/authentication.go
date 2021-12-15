package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	reqContext "github.com/Omelman/trucking-api/src/context"
	"github.com/Omelman/trucking-api/src/server/handlers"
	"github.com/Omelman/trucking-api/src/service"
)

// Auth - authenticate User by JWT token and add to context his ID, role etc.
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			authToken = r.Header.Get(handlers.AccessTokenHeader)
		)

		loginSes, err := service.Get().Validate(authToken)
		if err != nil {
			log.Error("Token is invalid ", authToken)
			handlers.SendEmptyResponse(w, http.StatusUnauthorized)
			return
		}

		ctx := reqContext.WithUserID(r.Context(), loginSes.UserID)
		ctx = reqContext.WithUserRole(ctx, loginSes.GetUserRole())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
