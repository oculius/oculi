package authn

import (
	"context"
)

type (
	service[UserDTO any, DAO UserDAO[UserDTO]] struct {
		repository Repository[UserDTO, DAO]
	}
)

func (b *service[UserDTO, UserDAO]) Login(ctx context.Context, identifier string, password string) (UserDAO, bool, error) {
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

func (b *service[UserDTO, UserDAO]) Register(ctx context.Context, user UserDAO) error {
	found, err := b.repository.LookupIdentifier(ctx, user.Identifier())
	if err != nil {
		return ErrAuthnFailedServer(err, nil)
	}
	if found {
		return ErrAuthnFailedUser(nil, nil)
	}
	return b.repository.InsertUser(ctx, user)
}

func NewService[UserDTO any, DAO UserDAO[UserDTO]](
	repo Repository[UserDTO, DAO],
) Core[UserDTO, DAO] {
	return &service[UserDTO, DAO]{repository: repo}
}
