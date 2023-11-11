package authtoken

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	Claims[V any] struct {
		jwt.RegisteredClaims
		Data       V      `json:"data"`
		Identifier string `json:"identifier"`
	}
)

func (c *Claims[T]) SetExpires(exp time.Time) {
	c.ExpiresAt = jwt.NewNumericDate(exp)
}

func (c *Claims[T]) SetTime(time time.Time) {
	c.IssuedAt = jwt.NewNumericDate(time)
	c.NotBefore = jwt.NewNumericDate(time)
}

func (c *Claims[V]) Credentials() V {
	return c.Data
}

var _ JwtClaimContract = &Claims[struct{}]{}
