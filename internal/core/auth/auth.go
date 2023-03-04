package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type Authenticator struct {
	secretKey string
}

func NewAuthenticator(secretKey string) *Authenticator {
	return &Authenticator{
		secretKey: secretKey,
	}
}

// GenerateToken godoc
// @Summary Generate jwt which lasts 5 minutes
// @Description Returns a jwt token
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {string} string ""
// @Router /token [get]
func (auth *Authenticator) GenerateToken(w http.ResponseWriter, _ *http.Request) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(auth.secretKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Error while generating token")
	}

	_ = json.NewEncoder(w).Encode(signedToken)
}

func (auth *Authenticator) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if ok := auth.validateToken(r); !ok {
			w.WriteHeader(http.StatusUnauthorized)

			return
		}
		next(w, r)
	}
}

func (auth *Authenticator) SetSecretKey(secretKey string) {
	auth.secretKey = secretKey
}

func (auth *Authenticator) validateToken(r *http.Request) bool {
	bearerToken := auth.parseToken(r)

	_, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(auth.secretKey), nil
	})

	return err == nil
}

func (auth *Authenticator) parseToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return ""
	}

	return strings.TrimPrefix(authHeader, "Bearer ")
}
