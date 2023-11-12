package di

import (
	"fmt"

	"github.com/pkg/errors"
	"go.uber.org/fx"
)

type (
	diErrorLogger struct{}
)

func (d diErrorLogger) HandleError(err error) {
	fmt.Printf("%+v\n", errors.WithStack(err))
}

var _ fx.ErrorHandler = diErrorLogger{}
