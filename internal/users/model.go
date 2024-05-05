package users

import (
	"rr-backend/ent/entgen"
	"rr-backend/lib/restmdl"
	"time"
)

// request
type RQUser struct {
	Id       string  `-`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Email    *string `json:"email"`

	rmd restmdl.RequestMetaData

	UpdatedAt *time.Time `json:"updatedAt"`
}

func (rq RQUser) MapToAggregate() User {
	var user User

	user.Id = rq.Id
	user.Name = rq.Name
	user.Password = rq.Password
	user.Email = rq.Email

	user.UpdatedAt = rq.UpdatedAt

	return user
}

// agregate
type User struct {
	Id       string
	Name     string
	Password string
	Email    *string

	UpdatedAt *time.Time
}

func (u *User) MapFromEnt(entModel *entgen.TblUSers) {
	u.Id = entModel.ID
	u.Name = entModel.UserName
	u.Password = entModel.Password
	u.Email = entModel.Email

	u.UpdatedAt = &entModel.UpdatedAt
}

// response
type RSUser struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Email *string `json:"email"`

	UpdatedAt *time.Time `json:"updatedAt"`
}

func (rs *RSUser) MapFromAggregate(user User) {
	rs.Id = user.Id
	rs.Name = user.Name
	rs.Email = user.Email

	rs.UpdatedAt = user.UpdatedAt
}
