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

func getUserQuery(entClient *entgen.Client, id string) *entgen.TblUSersQuery {
	return entClient.TblUSers.Query().Where(
		tblusers.IDEQ(id),
	)
}

func getUserQueryTx(tx *entgen.Tx, id string) *entgen.TblUSersQuery {
	return tx.TblUSers.Query().Where(
		tblusers.IDEQ(id),
	)
}

func getUpdateUserQuery(tx *entgen.Tx, rq User) *entgen.TblUSersUpdateOne {
	return tx.TblUSers.UpdateOneID(rq.Id).SetNillableEmail(rq.Email).SetNillablePassword(&rq.Password)
}

func getDeleteUserQuery(tx *entgen.Tx, id string) *entgen.TblUSersDeleteOne {
	return tx.TblUSers.DeleteOneID(id)
}
