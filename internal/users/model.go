package users

import "rr-backend/ent/entgen"

// request
type RQUser struct {
	Name     string  `json:name`
	Password string  `json:password`
	Email    *string `json:email`
}

func (rq RQUser) MapToAggregate() User {
	var user User

	user.Name = rq.Name
	user.Password = rq.Password
	user.Email = rq.Email

	return user
}

// agregate
type User struct {
	Id       string
	Name     string
	Password string
	Email    *string
}

func (u *User) MapFromEnt(entModel entgen.TblUSers) {
	u.Id = entModel.ID
	u.Name = entModel.UserName
	u.Password = entModel.Password
	u.Email = entModel.Email
}

// response
type RSUser struct {
	Id    string  `json:id`
	Name  string  `json:name`
	Email *string `json:email`
}

func (rs *RSUser) MapFromAggregate(user User) {
	rs.Id = user.Id
	rs.Name = user.Name
	rs.Email = user.Email
}
