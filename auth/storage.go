package auth

import (
	"context"
	"errors"
	"rr-backend/config"
	"rr-backend/ent/entgen"
)

type IAuthStorage interface {
	Login(username, password string) (RSToken, error)
}

type sAuthStorage struct {
	entClient *entgen.Client
}

// Login implements IAuthStorage.
func (s *sAuthStorage) Login(username string, password string) (RSToken, error) {
	exist, err := getUsersQuery(s.entClient, username, password).Exist(context.Background())
	if err != nil {
		return RSToken{}, errors.New(err.Error())
	}

	if !exist {
		return RSToken{}, errors.New("User Not Found")
	}

	entUser, err := getUsersQuery(s.entClient, username, password).Only(context.Background())
	if err != nil {
		return RSToken{}, errors.New(err.Error())
	}

	token, err := CreateToken(entUser.ID)
	if err != nil {
		return RSToken{}, errors.New(err.Error())
	}

	rsToken := RSToken{AccessToken: token}

	return rsToken, nil
}

func NewAuthStorage(config config.IAppConfig) IAuthStorage {
	return &sAuthStorage{
		entClient: config.GetENTClient(),
	}
}
