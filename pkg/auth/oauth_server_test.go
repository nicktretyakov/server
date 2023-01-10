package auth_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pascaldekloe/jwt"
	"golang.org/x/oauth2"
)

const (
	userPortalCode = 51232
	userPhone      = "7990881241"
	userEmail      = "user@example.com"
	userProfileID  = "850bbe96-5675-4990-9e6b-0cdb2012d273"
)

func mockOAuthServer() *httptest.Server {
	r := mux.NewRouter()

	r.Use(JSONMiddleware)

	oauthRouter := r.PathPrefix("/oauth").Subrouter()
	oauthRouter.HandleFunc("/token", tokenHandler).Methods("POST")

	return httptest.NewServer(r)
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	accessToken := signToken(
		userProfileID,
		map[string]interface{}{
			"phone":       userPhone,
			"portal_code": strconv.Itoa(userPortalCode),
			"email":       userEmail,
		},
		"e0q2c",
		false,
	)

	WriteJSON(w, http.StatusOK, oauth2.Token{
		AccessToken:  accessToken,
		TokenType:    "bearer",
		RefreshToken: "",
		Expiry:       time.Now().Add(time.Hour),
	})
}

func signToken(sub string, set map[string]interface{}, secret string, isExpired bool) string {
	now := time.Now()
	if isExpired {
		now = now.Add(-5 * time.Hour)
	}

	accessExpires := now.Add(time.Hour)

	claims := jwt.Claims{
		Registered: jwt.Registered{
			Issued:  jwt.NewNumericTime(now.Truncate(time.Second)),
			Expires: jwt.NewNumericTime(accessExpires.Truncate(time.Second)),
			Subject: sub,
		},
		Set: set,
	}

	accessToken, err := claims.HMACSign(jwt.HS512, []byte(secret))
	if err != nil {
		panic(err)
	}

	return string(accessToken)
}

func WriteJSON(w http.ResponseWriter, status int, body interface{}) {
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(true)
	encoder.SetIndent("", "")

	if err := encoder.Encode(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_ = encoder.Encode(map[string]interface{}{"error": err.Error()})

		return
	}
}

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
