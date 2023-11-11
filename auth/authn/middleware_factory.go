package authn

import (
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/server/oculi"
	"github.com/oculius/oculi/v2/server/oculi/token"
)

type (
	middlewareFactory[UserDTO any] struct {
		AuthTokenEngine DefaultAuthenticationTokenEngine[UserDTO]
		TokenFactory    TokenFactory
		CredentialsKey  string
	}
)

func NewMiddlewareFactory[UserDTO any](credentialsKey string,
	tokenSource token.TokenSource, tokenSourceKey string,
	authTokenEngine DefaultAuthenticationTokenEngine[UserDTO]) MiddlewareFactory {
	return &middlewareFactory[UserDTO]{
		AuthTokenEngine: authTokenEngine,
		TokenFactory: func(isReq bool) token.Token {
			return token.T(tokenSource, tokenSourceKey, token.String, isReq)
		},
		CredentialsKey: credentialsKey,
	}
}

func (a *middlewareFactory[UserDTO]) AuthRequired(next oculi.HandlerFunc) oculi.HandlerFunc {
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

func (a *middlewareFactory[UserDTO]) AuthOptional(next oculi.HandlerFunc) oculi.HandlerFunc {
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
