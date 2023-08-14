package authn

import (
	"github.com/oculius/oculi/v2/common/auth-token"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/rest-server/oculi"
	"github.com/oculius/oculi/v2/rest-server/oculi/token"
)

type (
	authenticationEngine[V any] struct {
		AuthTokenEngine DefaultAuthenticationTokenEngine[V]
		TokenFactory    TokenFactory
		CredentialsKey  string
	}

	Factory[V any] interface {
		AuthRequired(next oculi.HandlerFunc) oculi.HandlerFunc
		AuthOptional(next oculi.HandlerFunc) oculi.HandlerFunc
	}

	DefaultAuthenticationTokenEngine[V any] authtoken.Engine[*authtoken.Claims[V]]
	TokenFactory                            func(isReq bool) token.Token
)

func NewAuthenticationEngine[V any](credentialsKey string,
	tokenSource token.TokenSource, tokenSourceKey string,
	authTokenEngine DefaultAuthenticationTokenEngine[V]) Factory[V] {
	return &authenticationEngine[V]{
		AuthTokenEngine: authTokenEngine,
		TokenFactory: func(isReq bool) token.Token {
			return token.T(tokenSource, tokenSourceKey, token.String, isReq)
		},
		CredentialsKey: credentialsKey,
	}
}

func (a *authenticationEngine[V]) AuthRequired(next oculi.HandlerFunc) oculi.HandlerFunc {
	return func(ctx oculi.Context) error {
		tkn := a.TokenFactory(true)
		tokenMap, err := ctx.Lookup(tkn)
		if err != nil {
			return ctx.JSON(err.ResponseCode(), response.New(err))
		}

		tokenString, err := token.TokenValue[string](tokenMap[tkn.Key()])
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
		tkn := a.TokenFactory(false)
		tokenMap, err := ctx.Lookup(tkn)
		if err != nil {
			return next(ctx)
		}

		tokenString, err := token.TokenValue[string](tokenMap[tkn.Key()])
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
