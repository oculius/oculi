package logs

import (
	"io"
	"os"
	"time"

	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

type (
	zapLogger struct {
		instance *zap.SugaredLogger
		prefix   string
		level    log.Lvl
		out      io.Writer
	}

	LoggerOption struct {
		DevelopmentMode bool
		Level           log.Lvl
		Prefix          string
		Output          io.Writer
	}
)

var _ Logger = &zapLogger{}

func NewZap(logOption LoggerOption, options ...zap.Option) (Logger, error) {
	var (
		instance *zap.Logger
		err      error
	)

	if logOption.DevelopmentMode {
		instance, err = zap.NewDevelopment(options...)
	} else {
		instance, err = zap.NewProduction(options...)
	}
	if err != nil {
		return nil, err
	}

	return &zapLogger{
		instance: instance.Sugar(),
		level:    logOption.Level,
		prefix:   logOption.Prefix,
		out:      logOption.Output,
	}, nil
}

// NewZapDevelopment is a method for constructing new Zap Logger for development env
func NewZapDevelopment(options ...zap.Option) (Logger, error) {
	return NewZap(LoggerOption{true, log.DEBUG, "", os.Stdout}, options...)
}

// NewZapProduction is a method for constructing new Zap Logger for production env
func NewZapProduction(options ...zap.Option) (Logger, error) {
	return NewZap(LoggerOption{false, log.INFO, "", os.Stdout}, options...)
}

func (z *zapLogger) Output() io.Writer {
	return z.out
}

func (z *zapLogger) SetOutput(w io.Writer) {
	z.out = w
}

func (z *zapLogger) Prefix() string {
	return z.prefix
}

func (z *zapLogger) SetPrefix(p string) {
	z.prefix = p
}

func (z *zapLogger) Level() log.Lvl {
	return z.level
}

func (z *zapLogger) SetLevel(v log.Lvl) {
	z.level = v
}

func (z *zapLogger) SetHeader(_ string) {
}

func (z *zapLogger) isLogginOn() bool {
	return z.level != log.OFF
}

func (z *zapLogger) Print(args ...interface{}) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Info(args...)
}

func (z *zapLogger) Printf(format string, args ...interface{}) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Infof(format, args...)
}

func (z *zapLogger) Printj(j log.JSON) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Infof("%+v\n", j)
}

func (z *zapLogger) Debug(args ...interface{}) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Debug(args...)
}

func (z *zapLogger) Debugf(format string, args ...interface{}) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Debugf(format, args...)
}

func (z *zapLogger) Debugj(j log.JSON) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Debugf("%+v\n", j)
}

func (z *zapLogger) Info(args ...interface{}) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Info(args...)
}

func (z *zapLogger) Infof(format string, args ...interface{}) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Infof(format, args...)
}

func (z *zapLogger) Infoj(j log.JSON) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Infof("%+v\n", j)
}

func (z *zapLogger) Warn(args ...interface{}) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Warn(args...)
}

func (z *zapLogger) Warnf(format string, args ...interface{}) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Warnf(format, args...)
}

func (z *zapLogger) Warnj(j log.JSON) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Warnf("%+v\n", j)
}

func (z *zapLogger) Error(args ...interface{}) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Error(args...)
}

func (z *zapLogger) Errorf(format string, args ...interface{}) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Errorf(format, args...)
}

func (z *zapLogger) Errorj(j log.JSON) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Errorf("%+v\n", j)
}

func (z *zapLogger) Fatal(args ...interface{}) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Fatal(args...)
}

func (z *zapLogger) Fatalj(j log.JSON) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Fatalf("%+v\n", j)
}

func (z *zapLogger) Fatalf(format string, args ...interface{}) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Fatalf(format, args...)
}

func (z *zapLogger) Panic(args ...interface{}) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Panic(args...)
}

func (z *zapLogger) Panicj(j log.JSON) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Panicf("%+v\n", j)
}

func (z *zapLogger) Panicf(format string, args ...interface{}) {
	if !z.isLogginOn() {
		return
	}

	z.instance.Panicf(format, args...)
}

func (z *zapLogger) Instance() interface{} {
	return z.instance
}

func (z *zapLogger) setupBaseLogger(info Info) *zap.SugaredLogger {
	return z.instance.With("metadata", info.Metadata(), "timestamp", time.Now().UTC().Format(time.RFC3339Nano))
}

func (z *zapLogger) OPrint(info Info) {
	if !z.isLogginOn() {
		return
	}

	z.setupBaseLogger(info).Info(info.Message())
}

func (z *zapLogger) ODebug(info Info) {
	if !z.isLogginOn() {
		return
	}

	z.setupBaseLogger(info).Debug(info.Message())
}

func (z *zapLogger) OInfo(info Info) {
	if !z.isLogginOn() {
		return
	}

	z.setupBaseLogger(info).Info(info.Message())
}

func (z *zapLogger) OWarn(info Info) {
	if !z.isLogginOn() {
		return
	}

	z.setupBaseLogger(info).Warn(info.Message())
}

func (z *zapLogger) OError(info Info) {
	if !z.isLogginOn() {
		return
	}

	z.setupBaseLogger(info).Error(info.Message())
}

func (z *zapLogger) OFatal(info Info) {
	if !z.isLogginOn() {
		return
	}

	z.setupBaseLogger(info).Fatal(info.Message())
}

func (z *zapLogger) OPanic(info Info) {
	if !z.isLogginOn() {
		return
	}

	z.setupBaseLogger(info).Panic(info.Message())
}
