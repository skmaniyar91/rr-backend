package users

import (
	"errors"
	"rr-backend/config"
	"rr-backend/ent/entgen"
	"rr-backend/lib"
	"rr-backend/lib/restmdl"
)

type IUserStorage interface {
	CreateUser(rq User, rmd restmdl.RequestMetaData) (string, error)
	GetUser(id string) (User, error)
	UpdateUser(id string) error
	DeleteUser(id string) error
}

type sUserStorage struct {
	entClient *entgen.Client
}

func NewUserStorage(config config.IAppConfig) IUserStorage {
	return &sUserStorage{
		entClient: config.GetENTClient(),
	}
}

// CreateUser implements IUserStorage.
func (s *sUserStorage) CreateUser(rq User, rmd restmdl.RequestMetaData) (string, error) {
	var id string
	ctx := lib.GetContextWithRequestMetadata(rmd)

	err := lib.RunInEntTransaction(s.entClient, ctx, func(tx *entgen.Tx) error {
		// check if username already exists.
		exist, err := getUserTxQuery(tx, rq.Name).Exist(ctx)
		if err != nil {
			return err
		}

		if exist {
			return errors.New("username already exist")
		}

		entUser, err := getCreateUserQuery(tx, rq).Save(ctx)
		if err != nil {
			return err
		}
		id = entUser.ID

		return nil
	})

	if err != nil {
		return "", err
	}

	return id, err
}

// DeleteUser implements IUserStorage.
func (s *sUserStorage) DeleteUser(id string) error {
	panic("unimplemented")
}

// GetUser implements IUserStorage.
func (s *sUserStorage) GetUser(id string) (User, error) {
	panic("unimplemented")
}

// UpdateUser implements IUserStorage.
func (s *sUserStorage) UpdateUser(id string) error {
	panic("unimplemented")
}
