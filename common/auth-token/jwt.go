package authtoken

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type (
	Claims[V any] struct {
		jwt.RegisteredClaims
		Data       V      `json:"data"`
		Identifier string `json:"identifier"`
	}

	jwtEngine[T ClaimContract] struct {
		key      []byte
		alg      Algorithm
		contract ClaimContract
	}

	Algorithm string
)

func (j *jwtEngine[T]) Contract() ClaimContract {
	return j.contract
}

func (c *Claims[T]) SetExpires(exp time.Time) {
	c.ExpiresAt = jwt.NewNumericDate(exp)
}

func (c *Claims[T]) SetTime(time time.Time) {
	c.IssuedAt = jwt.NewNumericDate(time)
	c.NotBefore = jwt.NewNumericDate(time)
}

var _ ClaimContract = &Claims[struct{}]{}

func (j *jwtEngine[T]) Encode(claim T, exp time.Duration) (string, error) {
	now := time.Now()
	claim.SetExpires(now.Add(exp))
	claim.SetTime(now)

	newToken := jwt.NewWithClaims(jwt.GetSigningMethod(j.alg.String()), claim)

	signedToken, err := newToken.SignedString(j.key)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (j *jwtEngine[T]) Decode(tokenString string) (T, error) {
	var emptyclaim T
	token, err := j.getToken(tokenString, j.contract)
	if err != nil {
		return emptyclaim, ErrInvalidToken(err, err.Error())
	}

	if claims, ok := token.Claims.(T); ok && token.Valid {
		return claims, nil
	} else if !token.Valid {
		return emptyclaim, ErrInvalidToken(nil, nil)
	}

	return emptyclaim, ErrInvalidToken(err, "unknown error")
}

func (j *jwtEngine[T]) Validate(tokenString string) bool {
	token, err := j.getToken(tokenString, j.contract)
	if err != nil {
		return false
	}

	if _, ok := token.Claims.(T); ok && token.Valid {
		return true
	}
	return false
}

func (a Algorithm) String() string {
	return string(a)
}

func (j *jwtEngine[T]) getToken(token string, obj ClaimContract) (*jwt.Token, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token, obj,
		func(t *jwt.Token) (interface{}, error) {
			return j.key, nil
		},
	)
	if err != nil {
		return nil, err
	}

	if tokenClaims == nil || tokenClaims.Claims == nil {
		return nil, errors.New("failed to parse token")
	}
	return tokenClaims, nil
}

func (c *Claims[V]) Credentials() V {
	return c.Data
}

func NewJWT[V any](key string, alg Algorithm) Engine[*Claims[V]] {
	return &jwtEngine[*Claims[V]]{
		key:      []byte(key),
		alg:      alg,
		contract: &Claims[V]{}}
}
