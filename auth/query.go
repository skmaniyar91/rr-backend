package auth

import (
	"rr-backend/ent/entgen"
	"rr-backend/ent/entgen/tblusers"
)

func getUsersQuery(entClient *entgen.Client, username, password string) *entgen.TblUSersQuery {
	return entClient.TblUSers.Query().
		Where(tblusers.UserNameEQ(username), tblusers.PasswordEQ(password))
}
