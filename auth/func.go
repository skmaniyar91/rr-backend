package auth

import (
	"errors"
	"rr-backend/config/env"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	})

	tokenStr, err := token.SignedString([]byte(env.GetJWTKey()))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateToken(tokenStr string) (*Claims, error) {
	var claims *Claims
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(env.GetJWTKey()), nil
	})

	if err != nil {
		return &Claims{}, errors.New("invalid token")
	}

	if !token.Valid {
		return &Claims{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)

	if ok {
		return claims, nil
	}

	return claims, err
}
