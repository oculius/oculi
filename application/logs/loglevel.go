package logs

import (
	"strings"

	"github.com/labstack/gommon/log"
)

func GetLoggerLevel(level string) log.Lvl {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return log.DEBUG
	case "INFO":
		return log.INFO
	case "WARN":
		return log.WARN
	case "ERROR":
		return log.ERROR
	default:
		return log.INFO
	}
}
