package authn

import (
	"context"
	errext "github.com/oculius/oculi/v2/common/error-extension"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type (
	user struct {
		Username string
		Name     string
		Password string
	}

	userDTO struct {
		Username string `json:"username"`
		Name     string `json:"name"`
	}

	loginDTO struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	registerDTO struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	basicTestRepo struct {
		dataTable map[string]user
	}
)

func newtestrepo() BasicRepository[userDTO, user] {
	return &basicTestRepo{
		dataTable: map[string]user{},
	}
}

func (b *basicTestRepo) GetByIdentifier(ctx context.Context, identifier string) (user, error) {
	if result, ok := b.dataTable[identifier]; !ok {
		return result, errors.New("user not found")
	} else {
		return result, nil
	}
}

func (b *basicTestRepo) VerifyPassword(password string, u user) error {
	if password == u.Password {
		return nil
	}
	return errors.New("wrong password")
}

func (b *basicTestRepo) LookupIdentifier(ctx context.Context, identifier string) (bool, error) {
	_, ok := b.dataTable[identifier]
	return ok, nil
}

func (b *basicTestRepo) InsertUser(ctx context.Context, u user) error {
	id := u.Identifier()
	if _, ok := b.dataTable[id]; !ok {
		b.dataTable[id] = u
		return nil
	} else {
		return errors.New("user already exists")
	}
}

func (u user) DTO() userDTO {
	return userDTO{u.Username, u.Name}
}

func (u user) Identifier() string {
	return u.Username
}

var (
	_ UserDAO[userDTO]               = user{}
	_ BasicRepository[userDTO, user] = &basicTestRepo{}
)

func Test_BasicAuthentication(t *testing.T) {
	serviceFactory := func() Basic[userDTO, user] {
		return NewBasicService[userDTO, user](newtestrepo())
	}
	repo := newtestrepo()
	unknownErr := errors.New("unknown error")
	mockRepo := NewBasicRepository[userDTO, user](repo.InsertUser, repo.GetByIdentifier, repo.VerifyPassword,
		func(ctx context.Context, identifier string) (bool, error) {
			return false, unknownErr
		})
	mockRepo2 := NewBasicRepository[userDTO, user](repo.InsertUser, func(ctx context.Context, identifier string) (user, error) {
		return user{}, unknownErr
	}, repo.VerifyPassword, func(ctx context.Context, identifier string) (bool, error) {
		return true, nil
	})
	service1 := serviceFactory()
	service3 := NewBasicService(mockRepo)
	service4 := NewBasicService(mockRepo2)

	t.Run("register", func(tt *testing.T) {
		err := service1.Register(context.Background(), user{
			Name:     "asd",
			Username: "asd",
			Password: "1234",
		})
		assert.Nil(tt, err)

		err = service1.Register(context.Background(), user{
			Name:     "asd",
			Username: "asd",
			Password: "1234",
		})
		assert.Error(tt, err)

		err = service3.Register(context.Background(), user{
			Name:     "asd",
			Username: "asd",
			Password: "1234",
		})
		assert.Error(tt, err)
		aerr, ok := err.(errext.Error)
		assert.True(tt, ok)
		assert.Equal(tt, 500, aerr.ResponseCode())

		err = service1.Register(context.Background(), user{
			Name:     "asd2",
			Username: "asd2",
			Password: "5678",
		})
		assert.Nil(tt, err)
	})

	t.Run("login", func(tt *testing.T) {
		_, login, err := service3.Login(context.Background(), "asd", "123")
		assert.Error(tt, err)
		aerr, ok := err.(errext.Error)
		assert.True(tt, ok)
		assert.Equal(tt, 500, aerr.ResponseCode())
		assert.False(tt, login)

		_, login, err = service1.Login(context.Background(), "asd3", "1234")
		assert.Error(tt, err)
		aerr, ok = err.(errext.Error)
		assert.True(tt, ok)
		assert.Equal(tt, 400, aerr.ResponseCode())
		assert.False(tt, login)

		_, login, err = service4.Login(context.Background(), "asd", "1234")
		assert.Error(tt, err)
		aerr, ok = err.(errext.Error)
		assert.True(tt, ok)
		assert.Equal(tt, 500, aerr.ResponseCode())
		assert.False(tt, login)

		_, login, err = service1.Login(context.Background(), "asd2", "wrongpassword")
		assert.Error(tt, err)
		aerr, ok = err.(errext.Error)
		assert.True(tt, ok)
		assert.Equal(tt, 400, aerr.ResponseCode())
		assert.False(tt, login)

		_, login, err = service1.Login(context.Background(), "asd", "1234")
		assert.True(tt, login)
		assert.Nil(tt, err)
	})
}
