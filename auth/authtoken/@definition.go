package authtoken

import (
	"github.com/golang-jwt/jwt/v5"
	errext "github.com/oculius/oculi/v2/common/error-extension"
	"time"
)

type (
	JwtEngine[T JwtClaimContract] interface {
		Encode(claim T, exp time.Duration) (string, errext.HttpError)
		Decode(token string) (T, errext.HttpError)
		Validate(token string) bool
		Contract() JwtClaimContract
	}

	JwtClaimContract interface {
		jwt.Claims
		SetExpires(exp time.Time)
		SetTime(time time.Time)
	}
)
