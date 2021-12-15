package policy

import (
	"net/http"

	common "github.com/Omelman/trucking-api/src/server/handlers"

	reqContext "github.com/Omelman/trucking-api/src/context"
	"github.com/Omelman/trucking-api/src/models"
)

func Owner(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx      = r.Context()
			userRole = reqContext.GetUserRole(ctx)
		)

		switch userRole {
		case string(models.UserRoleOwner):
			next.ServeHTTP(w, r)
			return
		default:
			common.SendHTTPError(w, models.ErrForbidden)
			return
		}
	})
}
