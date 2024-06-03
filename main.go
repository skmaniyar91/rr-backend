package main

import (
	"rr-backend/cmd"
	_ "rr-backend/docs"
)

//	@title						RR-Back-End
//	@version					1.0
//	@description				This is RR-BackEnd
//	@securityDefinitions.apikey	JWTAuthentication
//	@in							header
//	@name						access token
//
//	@description				Type "Bearer" followed by a space and JWT token.
func main() {
	cmd.Start()
}
