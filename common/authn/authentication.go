package authn

import (
	"github.com/oculius/oculi/v2/common/auth-token"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/server/oculi"
	"github.com/oculius/oculi/v2/server/oculi/token"
)

type (
	authenticationEngine[UserDTO any] struct {
		AuthTokenEngine DefaultAuthenticationTokenEngine[UserDTO]
		TokenFactory    TokenFactory
		CredentialsKey  string
	}

	MiddlewareFactory interface {
		AuthRequired(next oculi.HandlerFunc) oculi.HandlerFunc
		AuthOptional(next oculi.HandlerFunc) oculi.HandlerFunc
	}

	DefaultAuthenticationTokenEngine[UserDTO any] authtoken.Engine[*authtoken.Claims[UserDTO]]
	TokenFactory                                  func(isReq bool) token.Token
)

func NewAuthenticationEngine[UserDTO any](credentialsKey string,
	tokenSource token.TokenSource, tokenSourceKey string,
	authTokenEngine DefaultAuthenticationTokenEngine[UserDTO]) MiddlewareFactory {
	return &authenticationEngine[UserDTO]{
		AuthTokenEngine: authTokenEngine,
		TokenFactory: func(isReq bool) token.Token {
			return token.T(tokenSource, tokenSourceKey, token.String, isReq)
		},
		CredentialsKey: credentialsKey,
	}
}

func (a *authenticationEngine[UserDTO]) AuthRequired(next oculi.HandlerFunc) oculi.HandlerFunc {
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

func (a *authenticationEngine[UserDTO]) AuthOptional(next oculi.HandlerFunc) oculi.HandlerFunc {
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
