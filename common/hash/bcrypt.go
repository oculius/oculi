package hash

import "golang.org/x/crypto/bcrypt"

type (
	bcryptHash struct {
		cost int
	}
)

const (
	BcryptMinCost     = bcrypt.MinCost
	BcryptMaxCost     = bcrypt.MaxCost
	BcryptDefaultCost = bcrypt.DefaultCost
)

func (b *bcryptHash) Hash(raw string) (string, error) {
	buff := []byte(raw)

	hash, err := bcrypt.GenerateFromPassword(buff, b.cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (b *bcryptHash) Verify(raw string, hashed string) error {
	buffHash := []byte(hashed)
	buffRaw := []byte(raw)

	err := bcrypt.CompareHashAndPassword(buffHash, buffRaw)
	if err != nil {
		return err
	}

	return nil
}

func NewBcrypt() Hash {
	return &bcryptHash{BcryptDefaultCost}
}

func NewBcryptWithCost(cost int) Hash {
	if cost > BcryptMaxCost {
		cost = BcryptMaxCost
	} else if cost < BcryptMinCost {
		cost = BcryptMinCost
	}
	return &bcryptHash{cost}
}
