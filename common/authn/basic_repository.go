package authn

import "context"

type (
	basicRepository[UserDTO any, DAO UserDAO[UserDTO]] struct {
		insertUser       InsertUser[DAO]
		getByIdentifier  GetByIdentifier[DAO]
		verifyPassword   VerifyPassword[UserDTO, DAO]
		lookupIdentifier LookupIdentifier[DAO]
	}
)

func (b *basicRepository[UserDTO, DAO]) GetByIdentifier(ctx context.Context, identifier string) (DAO, error) {
	return b.getByIdentifier(ctx, identifier)
}

func (b *basicRepository[UserDTO, DAO]) VerifyPassword(password string, user DAO) error {
	return b.verifyPassword(password, user)
}

func (b *basicRepository[UserDTO, DAO]) LookupIdentifier(ctx context.Context, identifier string) (bool, error) {
	return b.lookupIdentifier(ctx, identifier)
}

func (b *basicRepository[UserDTO, DAO]) InsertUser(ctx context.Context, user DAO) error {
	return b.insertUser(ctx, user)
}

func NewBasicRepository[UserDTO any, DAO UserDAO[UserDTO]](
	iuFn InsertUser[DAO],
	idFn GetByIdentifier[DAO],
	verifyFn VerifyPassword[UserDTO, DAO],
	lookupFn LookupIdentifier[DAO]) BasicRepository[UserDTO, DAO] {
	return &basicRepository[UserDTO, DAO]{
		insertUser:       iuFn,
		getByIdentifier:  idFn,
		verifyPassword:   verifyFn,
		lookupIdentifier: lookupFn,
	}
}
