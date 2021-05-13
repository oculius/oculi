package oculi

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi/common"
	conn "github.com/ravielze/oculi/common/connection"
	mm "github.com/ravielze/oculi/common/module"
	"gorm.io/gorm"
)

type Test struct {
	Abc string
}

type InitFunction func(db *gorm.DB, g *gin.Engine)

func New(appName string, initModule InitFunction, initMiddleware InitFunction) {
	devMode := common.DevMode()
	var mode string
	if devMode {
		mode = "development"
	} else {
		mode = "production"
	}
	fmt.Printf("| \u001b[44;1mOculi\u001b[0m | Starting server %s in %s mode...\n", appName, mode)
	fmt.Printf("| \u001b[44;1mOculi\u001b[0m | Connecting to database...\n")
	db := conn.ConnectDatabase(devMode)
	fmt.Printf("| \u001b[44;1mOculi\u001b[0m | Creating Gin Engine...\n")
	g := gin.Default()

	fmt.Printf("| \u001b[44;1mOculi\u001b[0m | Initiating middleware...\n")
	initMiddleware(db, g)
	// middleware.InstallCors(g, []string{"http://localhost:3000", "https://example.com"})
	// middleware.InstallDefaultLimiter(g)

	fmt.Printf("| \u001b[44;1mOculi\u001b[0m | Initiating module...\n")
	initModule(db, g)
	// mm.AddModule(essentials.NewModule(db, g))

	mm.ShowModule()
	g.Run()
}
