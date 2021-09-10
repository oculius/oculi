package request

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/ravielze/oculi/common/baseX/radix36"
	"github.com/ravielze/oculi/common/model/dto/auth"
	consts "github.com/ravielze/oculi/constant/errors"
	keyConsts "github.com/ravielze/oculi/constant/key"
	"github.com/ravielze/oculi/persistent/sql"
)

type (
	base struct {
		ctx               context.Context
		db                sql.API
		tx                sql.API
		errors            []error
		responseCode      int
		data              map[string]interface{}
		requestIdentifier auth.StandardCredentials
		beforeRollback    []func()
		beforeCommit      []func() error
		afterRollback     []func()
		afterCommit       []func()
	}
)

func (r *base) WithIdentifier(id auth.StandardCredentials) ReqContext {
	r.requestIdentifier = id
	return r
}

func (r *base) Identifier() auth.StandardCredentials {
	return r.requestIdentifier
}

func (r *base) Context() context.Context {
	return r.ctx
}

func (r *base) HasError() bool {
	return len(r.errors) > 0
}

func (r *base) AddError(responseCode int, err ...error) {
	if r.responseCode < 400 {
		r.responseCode = responseCode
	}
	r.errors = append(r.errors, err...)
}

func (r *base) SetResponseCode(code int) {
	r.responseCode = code
}

func (r *base) ResponseCode() int {
	return r.responseCode
}
func (r *base) Error() error {
	if len(r.errors) > 0 {
		return r.errors[0]
	}
	return nil
}

func (r *base) HasTransaction() bool {
	return r.tx != nil
}

func (r *base) Transaction() sql.API {
	if r.tx == nil {
		return r.db
	}
	return r.tx
}

func (r *base) NewTransaction() {
	r.tx = r.db.Begin()
	r.afterCommit = nil
	r.afterRollback = nil
	r.beforeCommit = nil
	r.beforeRollback = nil
}

func (r *base) CommitTransaction() error {
	if r.tx == nil {
		return nil
	}

	if r.beforeCommit != nil {
		for _, f := range r.beforeCommit {
			if err := f(); err != nil {
				r.RollbackTransaction()
				return err
			}
		}
	}
	r.tx.Commit()
	r.tx = nil
	if r.afterCommit != nil {
		for _, f := range r.afterCommit {
			f()
		}
	}
	return nil
}

func (r *base) RollbackTransaction() {
	if r.tx == nil {
		return
	}

	if r.beforeRollback != nil {
		for _, f := range r.beforeRollback {
			f()
		}
	}
	r.tx.Rollback()
	r.tx = nil
	if r.afterRollback != nil {
		for _, f := range r.afterRollback {
			f()
		}
	}
}

func (r *base) BeforeRollbackDo(f func()) {
	r.beforeRollback = append(r.beforeRollback, f)
}

func (r *base) BeforeCommitDo(f func() error) {
	r.beforeCommit = append(r.beforeCommit, f)
}

func (r *base) AfterRollbackDo(f func()) {
	r.afterRollback = append(r.afterRollback, f)
}

func (r *base) AfterCommitDo(f func()) {
	r.afterCommit = append(r.afterCommit, f)
}

func NewBase(db sql.API) ReqContext {
	result := &base{
		ctx:               context.Background(),
		db:                db,
		tx:                nil,
		errors:            make([]error, 0),
		responseCode:      200,
		data:              make(map[string]interface{}, 5),
		requestIdentifier: auth.StandardCredentials{Metadata: "unauthorized"},
		beforeRollback:    nil,
		beforeCommit:      nil,
		afterRollback:     nil,
		afterCommit:       nil,
	}
	result.Set("isTransformed", false)
	return result
}

func (r *base) WithContext(ctx context.Context) ReqContext {
	r.ctx = ctx
	return r
}

func (r *base) ParseString(key, value string) ReqContext {
	if !r.HasError() {
		r.data[key] = value
	}
	return r
}
func (r *base) ParseStringOrDefault(key, value, def string) ReqContext {
	if !r.HasError() {
		if len(value) == 0 || len(strings.TrimSpace(value)) == 0 {
			value = def
		}
		r.data[key] = value
	}
	return r
}

func (r *base) ParseUUID(key, value string) ReqContext {
	if !r.HasError() {
		if len(value) == 0 || len(strings.TrimSpace(value)) == 0 {
			r.AddError(http.StatusBadRequest, errors.New(consts.ErrRequestMissingValue+key))
		} else {
			uuidParsed := uuid.FromStringOrNil(value)
			if strings.EqualFold(value, "default") {
				r.data[key] = "default"
			} else if uuidParsed == uuid.Nil {
				r.AddError(http.StatusBadRequest, errors.New(consts.ErrRequestValueNotUUID+key))
			} else {
				r.data[key] = uuidParsed.String()
			}
		}
	}
	return r
}

func (r *base) Parse36(key, value string) ReqContext {
	if !r.HasError() {
		p := strings.ToUpper(value)
		if strings.EqualFold(p, "default") {
			r.data[key] = "default"
		} else if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			r.AddError(http.StatusBadRequest, errors.New(consts.ErrRequestMissingValue+key))
		} else {
			if data, err := radix36.NewRadix36(p); err != nil {
				r.AddError(http.StatusBadRequest, errors.New(consts.ErrRequestValueNotBase36+key))
			} else {
				r.data[key] = data.String()
			}
		}
	}
	return r
}

func (r *base) ParseUUID36(key, value string) ReqContext {
	if !r.HasError() {
		p := value
		if strings.EqualFold(p, "default") {
			r.data[key] = "default"
		} else if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			r.AddError(http.StatusBadRequest, errors.New(consts.ErrRequestMissingValue+key))
		} else {
			data, err := radix36.NewFromUUIDString(p)
			if err != nil {
				r.AddError(http.StatusBadRequest, errors.New(consts.ErrRequestValueNotUUID+key))
			} else {
				r.data[key] = data.String()
			}
		}
	}
	return r
}

func (r *base) Parse36UUID(key, value string) ReqContext {
	if !r.HasError() {

		p := value
		if strings.EqualFold(p, "default") {
			r.data[key] = "default"
		} else if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			r.AddError(http.StatusBadRequest, errors.New(consts.ErrRequestMissingValue+key))
		} else {
			if data, err := radix36.NewRadix36(p); err != nil {
				r.AddError(http.StatusBadRequest, errors.New(consts.ErrRequestValueNotBase36+key))
			} else {
				r.data[key] = data.ToUUID().String()
			}
		}
	}
	return r
}

func (r *base) ParseBoolean(key, value string, def bool) ReqContext {
	if !r.HasError() {

		q := value
		if (len(q) == 0 || len(strings.TrimSpace(q)) == 0) ||
			(q != strconv.FormatBool(false) && q != strconv.FormatBool(true)) {
			q = strconv.FormatBool(def)
		}

		p := (q == "true")
		r.data[key] = p
	}
	return r
}

//Not safe, for safety, use Get or Set instead.
func (r *base) Data() map[string]interface{} {
	if r.HasError() {
		return nil
	}
	return r.data
}

func (r *base) Get(key string) (interface{}, error) {
	isTransformed := r.data["isTransformed"].(bool)
	if strings.HasPrefix(key, keyConsts.EchoPrefixConstant) && !isTransformed {
		return nil, consts.ErrNeedDataTransformation
	}

	result, ok := r.data[key]
	if !ok {
		return nil, consts.ErrKeyNotFound
	}

	return result, nil
}
func (r *base) GetOrDefault(key string, val interface{}) interface{} {
	isTransformed := r.data["isTransformed"].(bool)
	if strings.HasPrefix(key, keyConsts.EchoPrefixConstant) && !isTransformed {
		return val
	}

	result, ok := r.data[key]
	if !ok {
		return val
	}

	return result
}

func (r *base) Set(key string, val interface{}) {
	r.data[key] = val
}
