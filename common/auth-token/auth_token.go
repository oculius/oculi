package authtoken

import (
	"github.com/golang-jwt/jwt/v5"
	errext "github.com/oculius/oculi/v2/common/error-extension"
	"net/http"
	"time"
)

type (
	Engine[T ClaimContract] interface {
		Encode(claim T, exp time.Duration) (string, errext.Error)
		Decode(token string) (T, errext.Error)
		Validate(token string) bool
		Contract() ClaimContract
	}

	ClaimContract interface {
		jwt.Claims
		SetExpires(exp time.Time)
		SetTime(time time.Time)
	}
)

var (
	ErrInvalidToken = errext.New("invalid token", http.StatusUnauthorized)
	ErrFailedToSign = errext.New("failed to sign", http.StatusInternalServerError)

	HS256 = Algorithm(jwt.SigningMethodHS256.Name)
	HS384 = Algorithm(jwt.SigningMethodHS384.Name)
	HS512 = Algorithm(jwt.SigningMethodHS512.Name)
)
