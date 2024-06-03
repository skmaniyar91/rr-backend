package auth

import "github.com/golang-jwt/jwt"

// jwt
type (
	Claims struct {
		UserId string
		jwt.StandardClaims
	}
	RSToken struct {
		AccessToken string `json:"accessToken"`
	}
)

// user

type (
	User struct {
		Id       string `db:"Id_ulid"`
		UserName string `db:"UserName"`
		Password string `db:Password`
	}
)
