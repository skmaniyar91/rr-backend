package auth

import (
	"context"
	"rr-backend/config"
	"rr-backend/ent/entgen"
	"rr-backend/errorx"
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
		return RSToken{}, errorx.WrapENTError("auth", err)
	}

	if !exist {
		return RSToken{}, errorx.NewUnProccessableEntity("auth", "User Not Found")
	}

	entUser, err := getUsersQuery(s.entClient, username, password).Only(context.Background())
	if err != nil {
		return RSToken{}, errorx.WrapENTError("auth", err)
	}

	token, err := CreateToken(entUser.ID)
	if err != nil {
		return RSToken{}, errorx.WrapENTError("auth", err)
	}

	rsToken := RSToken{AccessToken: token}

	return rsToken, nil
}

func NewAuthStorage(config config.IAppConfig) IAuthStorage {
	return &sAuthStorage{
		entClient: config.GetENTClient(),
	}
}
