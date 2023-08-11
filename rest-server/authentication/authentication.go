package authentication

import (
	authtoken "github.com/oculius/oculi/v2/common/auth-token"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/rest-server/oculi"
	"github.com/oculius/oculi/v2/rest-server/oculi/token"
)

type (
	authenticationEngine[V any] struct {
		AuthTokenEngine DefaultAuthenticationTokenEngine[V]
		Token           token.Token
		CredentialsKey  string
	}

	AuthenticationEngine[V any] interface {
		AuthRequired(next oculi.HandlerFunc) oculi.HandlerFunc
		AuthOptional(next oculi.HandlerFunc) oculi.HandlerFunc
	}

	DefaultAuthenticationTokenEngine[V any] authtoken.Engine[*authtoken.Claims[V]]
)

func NewAuthenticationEngine[V any](credentialsKey string, token token.Token,
	authTokenEngine DefaultAuthenticationTokenEngine[V]) AuthenticationEngine[V] {
	return &authenticationEngine[V]{
		AuthTokenEngine: authTokenEngine,
		Token:           token,
		CredentialsKey:  credentialsKey,
	}
}

func (a *authenticationEngine[V]) AuthRequired(next oculi.HandlerFunc) oculi.HandlerFunc {
	return func(ctx oculi.Context) error {
		tokenMap, err := ctx.Lookup(a.Token)
		if err != nil {
			return ctx.JSON(err.ResponseCode(), response.New(err))
		}

		tokenString, err := token.TokenValue[string](tokenMap[a.Token.Key()])
		if err != nil {
			return ctx.JSON(err.ResponseCode(), response.New(err))
		}

		creds, err := a.AuthTokenEngine.Decode(tokenString)
		if err != nil {
			return ctx.JSON(err.ResponseCode(), response.New(err))
		}

		ctx.Set(a.CredentialsKey, *creds)
		return next(ctx)
	}
}

func (a *authenticationEngine[V]) AuthOptional(next oculi.HandlerFunc) oculi.HandlerFunc {
	return func(ctx oculi.Context) error {
		tokenMap, err := ctx.Lookup(a.Token)
		if err != nil {
			return next(ctx)
		}

		tokenString, err := token.TokenValue[string](tokenMap[a.Token.Key()])
		if err != nil {
			return next(ctx)
		}

		creds, err := a.AuthTokenEngine.Decode(tokenString)
		if err != nil {
			return next(ctx)
		}

		ctx.Set(a.CredentialsKey, *creds)
		return next(ctx)
	}
}
