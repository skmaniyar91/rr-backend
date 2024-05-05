package users

import (
	"rr-backend/ent/entgen"
	"rr-backend/ent/entgen/tblusers"
)

func getCreateUserQuery(tx *entgen.Tx, rq User) *entgen.TblUSersCreate {
	return tx.TblUSers.Create().SetUserName(rq.Name).SetPassword(rq.Password).SetNillableEmail(rq.Email)
}

func getUserTxQuery(tx *entgen.Tx, username string) *entgen.TblUSersQuery {
	return tx.TblUSers.Query().Where(
		tblusers.UserNameEQ(username),
	)
}
