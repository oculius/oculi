package rest

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"github.com/oculius/oculi/v2/internal/constant"
	"net/http"
	"sort"
	"time"
)

type (
	printer struct {
		colorizer *color.Color
	}
)

var printerInstance = printer{colorizer: color.New()}

func (p printer) methodColor(method string) string {
	methodStr := fmt.Sprintf(" %-7s ", method)
	switch method {
	case http.MethodGet:
		return p.colorizer.BlueBg(methodStr, color.Blk)
	case http.MethodPost:
		return p.colorizer.CyanBg(methodStr, color.Blk)
	case http.MethodPut:
		return p.colorizer.YellowBg(methodStr, color.Blk)
	case http.MethodDelete:
		return p.colorizer.RedBg(methodStr, color.Blk)
	case http.MethodPatch:
		return p.colorizer.GreenBg(methodStr, color.Blk)
	case http.MethodHead:
		return p.colorizer.MagentaBg(methodStr, color.Blk)
	case http.MethodOptions:
		return p.colorizer.WhiteBg(methodStr, color.Blk)
	case http.MethodTrace:
		return p.colorizer.BlackBg(methodStr, color.Wht)
	default:
		return p.colorizer.Reset(methodStr)
	}
}

func (p printer) statusCodeColor(code int) string {
	codeStr := fmt.Sprintf(" %3d ", code)
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return p.colorizer.GreenBg(codeStr, color.Blk)
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return p.colorizer.WhiteBg(codeStr, color.Blk)
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return p.colorizer.YellowBg(codeStr, color.Blk)
	default:
		return p.colorizer.RedBg(codeStr, color.Blk)
	}
}

func (p printer) fmtRequest(ec echo.Context, start time.Time) string {
	now := time.Now()
	statusCode := ec.Response().Status
	method := ec.Request().Method
	latency := now.Sub(start)
	path := ec.Path()
	if ec.QueryString() != "" {
		path = path + "?" + ec.QueryString()
	}
	return fmt.Sprintf("%-30v | %s | %13v | %15s |%-7s %s\n",
		now.Format(time.RFC1123),
		p.statusCodeColor(statusCode),
		latency,
		ec.RealIP(),
		p.methodColor(method),
		path,
	)
}

const banner string = `
  ____           ___ 
 / __ \______ __/ (_)
/ /_/ / __/ // / / /
\____/\__/\_,_/_/_/ %s (%d)
Compact & intact Go web framework
----------------------------------
`

func (p printer) printRoutes(ec *echo.Echo) {
	colorizer := color.New()
	colorizer.SetOutput(ec.Logger.Output())

	fmt.Printf(banner, colorizer.Red("v"+constant.Version), constant.VersionNumber)
	routes := make(map[string][]string)
	for _, route := range ec.Routes() {
		routes[route.Method] = append(routes[route.Method], route.Path)
	}

	methodList := []string{http.MethodGet, http.MethodPost, http.MethodPatch,
		http.MethodPut, http.MethodDelete, http.MethodHead, http.MethodConnect,
		http.MethodOptions, http.MethodTrace}

	for method := range routes {
		sort.Strings(routes[method])
	}

	fmt.Printf("%-30v | Showing registered routes:\n",
		time.Now().Format(time.RFC1123))

	for _, method := range methodList {
		for _, route := range routes[method] {
			fmt.Printf("%-30v | %s %s\n",
				"",
				p.methodColor(method),
				route,
			)
		}
	}
}
