package users

import (
	"rr-backend/config"
	"rr-backend/lib/restmdl"
)

type IUSerService interface {
	CreateUser(rq RQUser) (RSUser, error)
	GetUser(id string) (RSUser, error)
	UpdateUser(rq RQUser) (RSUser, error)
	DeleteUser(id string, rmd restmdl.RequestMetaData) error
}

type sUserService struct {
	userStorage IUserStorage
}

func NewUserService(config config.IAppConfig) IUSerService {
	return &sUserService{
		userStorage: NewUserStorage(config),
	}
}

// CreateUser implements IUSerService.
func (s *sUserService) CreateUser(rq RQUser) (RSUser, error) {
	user := rq.MapToAggregate()

	id, err := s.userStorage.CreateUser(user, rq.rmd)
	if err != nil {
		return RSUser{}, err
	}

	rs, err := s.GetUser(id)
	if err != nil {
		return RSUser{}, err
	}

	return rs, nil
}

// DeleteUser implements IUSerService.
func (s *sUserService) DeleteUser(id string, rmd restmdl.RequestMetaData) error {
	err := s.userStorage.DeleteUser(id, rmd)
	if err != nil {
		return err
	}

	return nil
}

// GetUser implements IUSerService.
func (s *sUserService) GetUser(id string) (RSUser, error) {
	user, err := s.userStorage.GetUser(id)
	if err != nil {
		return RSUser{}, err
	}

	var rs RSUser
	rs.MapFromAggregate(user)

	return rs, nil
}

// UpdateUser implements IUSerService.
func (s *sUserService) UpdateUser(rq RQUser) (RSUser, error) {
	user := rq.MapToAggregate()

	err := s.userStorage.UpdateUser(user, rq.rmd)
	if err != nil {
		return RSUser{}, err
	}

	rs, err := s.GetUser(rq.Id)
	if err != nil {
		return RSUser{}, err
	}

	return rs, nil
}
