package authtoken

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/oculius/oculi/v2/common/error-extension"
)

type (
	jwtEngine[T JwtClaimContract] struct {
		key      []byte
		alg      Algorithm
		contract JwtClaimContract
	}
)

func (j *jwtEngine[T]) Contract() JwtClaimContract {
	return j.contract
}

func (j *jwtEngine[T]) Encode(claim T, exp time.Duration) (string, errext.HttpError) {
	now := time.Now()
	claim.SetExpires(now.Add(exp))
	claim.SetTime(now)

	newToken := jwt.NewWithClaims(jwt.GetSigningMethod(j.alg.String()), claim)

	signedToken, err := newToken.SignedString(j.key)
	if err != nil {
		return "", ErrFailedToSign(err, nil)
	}
	return signedToken, nil
}

func (j *jwtEngine[T]) Decode(tokenString string) (T, errext.HttpError) {
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

func (j *jwtEngine[T]) getToken(token string, obj JwtClaimContract) (*jwt.Token, error) {
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

func NewJWT[V any](key string, alg Algorithm) JwtEngine[*Claims[V]] {
	return &jwtEngine[*Claims[V]]{
		key:      []byte(key),
		alg:      alg,
		contract: &Claims[V]{}}
}
