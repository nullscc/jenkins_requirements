package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nullscc/jenkins_requirements/middlewares"
	"github.com/nullscc/jenkins_requirements/routers/api"
	"os"
)

var (
	Router *gin.Engine
)

func init() {
	if len(os.Args) > 1 {
		return
	}
	// using default route with middlewares of Logger and Recovery enabled
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"token", "Content-Type"}
	r.Use(cors.New(config))

	apiG := r.Group("api")
	apiG.Use(middlewares.JWT())
	{
		apiG.GET("test", api.Test)
		apiG.POST("requirement", api.Add)
		apiG.GET("requirement", api.Get)
	}

	Router = r
}

func Run(addr ...string) {
	Router.Run(addr...)
}
