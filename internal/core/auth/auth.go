package auth

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"net/http"
	"strings"
	"time"
)

type Authenticator struct {
	secretKey string
}

func NewAuthenticator(secretKey string) *Authenticator {
	return &Authenticator{
		secretKey: secretKey,
	}
}

func (authenticator *Authenticator) GenerateToken(w http.ResponseWriter, _ *http.Request) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(authenticator.secretKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Error while generating token")
	}

	_ = json.NewEncoder(w).Encode(signedToken)
}

func (authenticator *Authenticator) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if ok := authenticator.validateToken(r); !ok {
			w.WriteHeader(http.StatusUnauthorized)

			return
		}
		next(w, r)
	}
}

func (authenticator *Authenticator) SetSecretKey(secretKey string) {
	authenticator.secretKey = secretKey
}

func (authenticator *Authenticator) validateToken(r *http.Request) bool {
	bearerToken := authenticator.parseToken(r)

	_, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(authenticator.secretKey), nil
	})
	if err != nil {
		return false
	}

	return true
}

func (authenticator *Authenticator) parseToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return ""
	}

	return strings.TrimPrefix(authHeader, "Bearer ")
}
