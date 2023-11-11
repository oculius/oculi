package authn

import "context"

type (
	repository[UserDTO any, DAO UserDAO[UserDTO]] struct {
		insertUser       InsertUser[DAO]
		getByIdentifier  GetByIdentifier[DAO]
		verifyPassword   VerifyPassword[UserDTO, DAO]
		lookupIdentifier LookupIdentifier[DAO]
	}
)

func (b *repository[UserDTO, DAO]) GetByIdentifier(
	ctx context.Context, identifier string) (DAO, error) {
	return b.getByIdentifier(ctx, identifier)
}

func (b *repository[UserDTO, DAO]) VerifyPassword(password string, user DAO) error {
	return b.verifyPassword(password, user)
}

func (b *repository[UserDTO, DAO]) LookupIdentifier(ctx context.Context, identifier string) (bool, error) {
	return b.lookupIdentifier(ctx, identifier)
}

func (b *repository[UserDTO, DAO]) InsertUser(ctx context.Context, user DAO) error {
	return b.insertUser(ctx, user)
}

func NewRepository[
	UserDTO any,
	DAO UserDAO[UserDTO],
](
	insertFn InsertUser[DAO],
	idFn GetByIdentifier[DAO],
	verifyFn VerifyPassword[UserDTO, DAO],
	lookupFn LookupIdentifier[DAO]) Repository[UserDTO, DAO] {
	return &repository[UserDTO, DAO]{
		insertUser:       insertFn,
		getByIdentifier:  idFn,
		verifyPassword:   verifyFn,
		lookupIdentifier: lookupFn,
	}
}
