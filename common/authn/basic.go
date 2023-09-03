package authn

import (
	"context"
	"github.com/oculius/oculi/v2/rest-server"
)

type (
	Basic[UserDTO any, DAO UserDAO[UserDTO]] interface {
		Login(ctx context.Context, identifier string, password string) (DAO, bool, error)
		Register(ctx context.Context, user DAO) error
	}

	UserDAO[DTO any] interface {
		DTO() DTO
		Identifier() string
	}

	BasicRepository[UserDTO any, DAO UserDAO[UserDTO]] interface {
		GetByIdentifier(ctx context.Context, identifier string) (DAO, error)
		VerifyPassword(password string, user DAO) error
		LookupIdentifier(ctx context.Context, identifier string) (bool, error)
		InsertUser(ctx context.Context, user DAO) error
	}

	BasicRestModule interface {
		rest.Module

		Login(ctx context.Context) error
		WhoAmI(ctx context.Context) error
		Register(ctx context.Context) error
	}

	GetByIdentifier[DAO any]                  func(ctx context.Context, identifier string) (DAO, error)
	VerifyPassword[DTO any, DAO UserDAO[DTO]] func(password string, user DAO) error
	LookupIdentifier[DAO any]                 func(ctx context.Context, identifier string) (bool, error)
	InsertUser[DAO any]                       func(ctx context.Context, user DAO) error
)
