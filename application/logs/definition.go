package logs

import (
	"io"

	"github.com/labstack/gommon/log"
)

type (
	Info interface {
		Message() string
		Metadata() map[string]any
		EditMetadata(key string, value any)
		RemoveMetadata(key string)
		ClearMetadata()
	}

	Metadata interface {
		Apply(si Info)
	}

	OculiLogger interface {
		// OPrint stands for Oculi Print. A method to print Info as info.
		OPrint(info Info)
		// ODebug stands for Oculi Debug. A method to print Info as debug.
		ODebug(info Info)
		// OInfo stands for Oculi Info. A method to print Info as info.
		OInfo(info Info)
		// OWarn stands for Oculi Warn. A method to print Info as warn.
		OWarn(info Info)
		// OError stands for Oculi Error. A method to print Info as error.
		OError(info Info)
		// OFatal stands for Oculi Fatal. A method to print Info as fatal.
		OFatal(info Info)
		// OPanic stands for Oculi Panic. A method to print Info as panic.
		OPanic(info Info)
	}

	Logger interface {
		OculiLogger
		Output() io.Writer
		SetOutput(w io.Writer)
		Prefix() string
		SetPrefix(p string)
		Level() log.Lvl
		SetLevel(v log.Lvl)
		SetHeader(h string)
		With(args ...any) Logger
		Print(i ...interface{})
		Printf(format string, args ...interface{})
		Printj(j log.JSON)
		Debug(i ...interface{})
		Debugf(format string, args ...interface{})
		Debugj(j log.JSON)
		Info(i ...interface{})
		Infof(format string, args ...interface{})
		Infoj(j log.JSON)
		Warn(i ...interface{})
		Warnf(format string, args ...interface{})
		Warnj(j log.JSON)
		Error(i ...interface{})
		Errorf(format string, args ...interface{})
		Errorj(j log.JSON)
		Fatal(i ...interface{})
		Fatalj(j log.JSON)
		Fatalf(format string, args ...interface{})
		Panic(i ...interface{})
		Panicj(j log.JSON)
		Panicf(format string, args ...interface{})
		Instance() interface{}
	}
)
