package authtoken

import "github.com/golang-jwt/jwt/v5"

type (
	Algorithm string
)

var (
	HS256 = Algorithm(jwt.SigningMethodHS256.Name)
	HS384 = Algorithm(jwt.SigningMethodHS384.Name)
	HS512 = Algorithm(jwt.SigningMethodHS512.Name)
)

func (a Algorithm) String() string {
	return string(a)
}
