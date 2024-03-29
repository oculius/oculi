package response

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	oculiContext "github.com/ravielze/oculi/context"
	"github.com/ravielze/oculi/errors"
	"github.com/ravielze/oculi/request"
	oculiValidator "github.com/ravielze/oculi/validator"
)

type (
	standardResponse struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}

	errorResponse struct {
		Error   string      `json:"error"`
		Details interface{} `json:"details"`
	}

	errorField struct {
		Field   string `json:"field"`
		Reason  string `json:"reason"`
		Message string `json:"message"`
	}

	errorHttp struct {
		Internal string      `json:"internal"`
		Message  interface{} `json:"error"`
	}

	responder struct {
		validator     oculiValidator.Validator
		isDevelopment bool
	}

	Responder interface {
		NewJSONResponse(ctx *oculiContext.Context, req request.ReqContext, data interface{}) error
	}

	Response interface{}
)

func New(validator oculiValidator.Validator, isDevelopment bool) Responder {
	return &responder{
		validator:     validator,
		isDevelopment: isDevelopment,
	}
}
func (r *responder) NewJSONResponse(ctx *oculiContext.Context, req request.ReqContext, data interface{}) error {
	if req != nil {
		ctx.Merge(req)
	}
	return r.newJSON(ctx, data)
}

func (r *responder) newJSON(ctx *oculiContext.Context, data interface{}) error {
	var resp Response
	if ctx.ResponseCode() >= 400 || ctx.HasError() {
		resp = standardResponse{
			Code: ctx.ResponseCode(),
			Data: r.handleError(ctx.ResponseCode(), ctx.Errors()),
		}
	} else if data == nil {
		resp = standardResponse{
			Code: ctx.ResponseCode(),
			Data: "success",
		}
	} else {
		resp = standardResponse{
			Code: ctx.ResponseCode(),
			Data: data,
		}
	}
	if r.isDevelopment {
		return ctx.JSONPretty(ctx.ResponseCode(), resp, " ")
	}
	return ctx.JSON(ctx.ResponseCode(), resp)
}

func (r *responder) handleError(responseCode int, data []error) errorResponse {
	if data == nil {
		return errorResponse{
			Error:   "unknown error",
			Details: []interface{}{},
		}
	}
	msg, errfields := r.buildErrors(responseCode, data)
	if errfields == nil {
		return errorResponse{
			Error:   msg,
			Details: []interface{}{},
		}
	}
	return errorResponse{
		Error:   msg,
		Details: errfields,
	}
}

func (r *responder) buildErrors(responseCode int, data []error) (string, interface{}) {
	derr, ok := data[0].(errors.DetailedErrors)
	if ok {
		details := errors.Details(derr)
		if len(details) == 1 {
			return derr.Error(), errors.Details(derr)[0]
		}
		return derr.Error(), errors.Details(derr)
	}

	if responseCode != http.StatusUnprocessableEntity {
		return data[0].Error(), nil
	}
	err, ok := data[0].(validator.ValidationErrors)
	if ok {
		errors := make([]errorField, len(err))
		for i := range errors {
			errors[i] = errorField{
				Field:   err[i].Namespace(),
				Reason:  err[i].Tag(),
				Message: err[i].Translate(*r.validator.Translator()),
			}
		}
		return "validation error", errors
	}

	d, ok := data[0].(*echo.HTTPError)
	if ok {
		internal := ""
		if d.Internal != nil {
			internal = d.Internal.Error()
		}
		return "json body is unparseable", errorHttp{
			Internal: internal,
			Message:  d.Message,
		}
	}

	return data[0].Error(), nil
}
