package auth

import "github.com/golang-jwt/jwt"

type (
	Claims struct {
		UserId string
		jwt.StandardClaims
	}
)

type RSToken struct {
	AccessToken string `json:accessToken`
}
