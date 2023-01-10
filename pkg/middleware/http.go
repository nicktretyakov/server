package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"be/internal/model"
	"be/pkg/auth"
	"be/pkg/response"
	"be/pkg/xerror"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"} //nolint:gochecknoglobals

func AuthMiddleware(authApp auth.IAuth) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, _, xerr := userFromBearerToken(authApp, r)
			if xerr != nil {
				response.WriteJSON(w, http.StatusUnauthorized, xerr)
				return
			}

			handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), &userCtxKey, user)))
		})
	}
}

func userFromBearerToken(authApp auth.IAuth, r *http.Request) (*model.User, *model.Session, *xerror.Error) {
	bearer := r.Header.Get("authorization")

	if len(bearer) < 7 || !strings.EqualFold(bearer[0:6], "bearer") {
		return nil, nil, xerror.New(xerror.PermissionDenied, "invalid token")
	}

	token := bearer[7:]

	return authApp.ParseToken(r.Context(), token)
}

func UserFromContext(ctx context.Context) *model.User {
	var (
		userPtr *model.User
		ok      bool
	)

	userValue := ctx.Value(&userCtxKey)
	if userPtr, ok = userValue.(*model.User); userValue == nil || !ok {
		return nil
	}

	return userPtr
}
