package request

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/common/model/dto/auth"
	"github.com/ravielze/oculi/persistent/sql"
)

type (
	ReqContext interface {
		WithContext(ctx context.Context) ReqContext
		Context() context.Context
		HasError() bool
		AddError(responseCode int, err ...error)

		SetResponseCode(code int)
		ResponseCode() int
		Error() error

		HasTransaction() bool
		Transaction() sql.API
		NewTransaction()
		CommitTransaction() error
		RollbackTransaction()
		BeforeRollbackDo(f func())
		BeforeCommitDo(f func() error)
		AfterRollbackDo(f func())
		AfterCommitDo(f func())

		ParseString(key, value string) ReqContext
		ParseStringOrDefault(key, value, def string) ReqContext
		ParseUUID(key, value string) ReqContext
		Parse36(key, value string) ReqContext
		ParseUUID36(key, value string) ReqContext
		Parse36UUID(key, value string) ReqContext
		ParseBoolean(key, value string, def bool) ReqContext

		Get(key string) (interface{}, error)
		GetOrDefault(key string, def interface{}) interface{}
		Set(key string, val interface{})

		Identifier() auth.StandardCredentials
		WithIdentifier(id auth.StandardCredentials) ReqContext
	}

	//TODO NotImplemented
	NonEchoContext interface {
		BindValidate(obj interface{})
	}

	EchoReqContext interface {
		ReqContext

		Echo() echo.Context

		Param(param string) EchoReqContext
		ParamUUID(param string) EchoReqContext
		Param36(param string) EchoReqContext
		ParamUUID36(param string) EchoReqContext
		Param36UUID(param string) EchoReqContext

		// Get query with string value and set it to default if it's empty
		Query(query, def string) EchoReqContext

		// Get query with boolean value
		QueryBoolean(query string, def bool) EchoReqContext

		// Transfer echo store data to request base data
		Transform() ReqContext
	}
)
