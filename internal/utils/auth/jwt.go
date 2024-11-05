package auth

import (
	"ecommerce/global"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type PayloadClaim struct {
	jwt.StandardClaims
}

func CrateToken(uuidToken string) (string, error) {
	// 1. set time expiration
	timeEx := global.Config.JWT.JWT_EXPIRATION
	if timeEx == "" {
		timeEx = "1h"
	}

	expiration, err := time.ParseDuration(timeEx)
	if err != nil {
		return "", err
	}

	now := time.Now()
	expiresAt := now.Add(expiration)

	return GenerateToken(&PayloadClaim{
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "showdevgo",
			Subject:   uuidToken,
		},
	})
}
func GenerateToken(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.JWT.API_SECRET_KEY))
}