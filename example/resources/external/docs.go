package external

import (
	_ "embed"

	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/docs"
	"github.com/ravielze/oculi/example/config"
)

//go:embed docs.json
var swaggerJSON string

func NewDocs(ec *echo.Echo, config *config.Env) docs.Documentation {
	if config.IsProduction() {
		return nil
	}
	docs.SetData(swaggerJSON)
	return docs.New(ec, config.ServiceName, config.ServiceHost)
}
