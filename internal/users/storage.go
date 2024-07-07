package users

import (
	"context"
	"errors"
	"rr-backend/config"
	"rr-backend/ent/entgen"
	"rr-backend/errorx"
	"rr-backend/lib"
	"rr-backend/lib/restmdl"
)

type IUserStorage interface {
	CreateUser(rq User, rmd restmdl.RequestMetaData) (string, error)
	GetUser(id string) (User, error)
	UpdateUser(rq User, rmd restmdl.RequestMetaData) error
	DeleteUser(id string, rmd restmdl.RequestMetaData) error
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
			return errorx.WrapENTError(Domain, err)
		}

		if exist {
			return errorx.NewUnProccessableEntity(Domain, "User name already exists")
		}

		entUser, err := getCreateUserQuery(tx, rq).Save(ctx)
		if err != nil {
			return errorx.WrapENTError(Domain, err)
		}
		id = entUser.ID

		return nil
	})

	if err != nil {
		return "", err
	}

	return id, nil
}

// DeleteUser implements IUserStorage.
func (s *sUserStorage) DeleteUser(id string, rmd restmdl.RequestMetaData) error {
	ctx := lib.GetContextWithRequestMetadata(rmd)

	err := lib.RunInEntTransaction(s.entClient, ctx, func(tx *entgen.Tx) error {
		// check if username already exists.
		err := getDeleteUserQuery(tx, id).Exec(ctx)
		if err != nil {
			return errorx.WrapENTError(Domain, err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// GetUser implements IUserStorage.
func (s *sUserStorage) GetUser(id string) (User, error) {
	entUser, err := getUserQuery(s.entClient, id).Only(context.Background())
	if err != nil {
		return User{}, errorx.WrapENTError(Domain, err)
	}

	var user User
	user.MapFromEnt(entUser)

	return user, nil
}

// UpdateUser implements IUserStorage.
func (s *sUserStorage) UpdateUser(rq User, rmd restmdl.RequestMetaData) error {
	ctx := lib.GetContextWithRequestMetadata(rmd)

	err := lib.RunInEntTransaction(s.entClient, ctx, func(tx *entgen.Tx) error {

		entUser, err := getUserQueryTx(tx, rq.Id).Only(ctx)
		if err != nil {
			return errorx.WrapENTError(Domain, err)
		}

		if !rq.UpdatedAt.Equal(entUser.UpdatedAt) {
			return errors.New("steal data")
		}

		err = getUpdateUserQuery(tx, rq).Exec(ctx)
		if err != nil {
			return errorx.WrapENTError(Domain, err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return err
}
