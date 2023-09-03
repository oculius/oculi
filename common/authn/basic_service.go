package authn

import (
	"context"
	errext "github.com/oculius/oculi/v2/common/error-extension"
	"net/http"
)

type (
	basicService[UserDTO any, DAO UserDAO[UserDTO]] struct {
		repository BasicRepository[UserDTO, DAO]
	}
)

var (
	ErrAuthnFailedServer = errext.New("failed, server error", http.StatusInternalServerError)
	ErrAuthnFailedUser   = errext.New("failed, user error", http.StatusBadRequest)
)

func (b *basicService[UserDTO, UserDAO]) Login(ctx context.Context, identifier string, password string) (UserDAO, bool, error) {
	var empty UserDAO
	found, err := b.repository.LookupIdentifier(ctx, identifier)
	if err != nil {
		return empty, false, ErrAuthnFailedServer(err, nil)
	}
	if !found {
		return empty, false, ErrAuthnFailedUser(nil, nil)
	}

	user, err := b.repository.GetByIdentifier(ctx, identifier)
	if err != nil {
		return empty, false, ErrAuthnFailedServer(err, nil)
	}

	err = b.repository.VerifyPassword(password, user)
	if err != nil {
		return empty, false, ErrAuthnFailedUser(err, nil)
	}
	return user, true, nil
}

func (b *basicService[UserDTO, UserDAO]) Register(ctx context.Context, user UserDAO) error {
	found, err := b.repository.LookupIdentifier(ctx, user.Identifier())
	if err != nil {
		return ErrAuthnFailedServer(err, nil)
	}
	if found {
		return ErrAuthnFailedUser(nil, nil)
	}
	return b.repository.InsertUser(ctx, user)
}

func NewBasicService[UserDTO any, DAO UserDAO[UserDTO]](
	repo BasicRepository[UserDTO, DAO],
) Basic[UserDTO, DAO] {
	return &basicService[UserDTO, DAO]{repository: repo}
}
