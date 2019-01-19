package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/nullscc/jenkins_requirements/utils"
	"os"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		docker_token := os.Getenv("REQUIREMENTS_TOKEN")
		if token != docker_token {
			utils.JsonRetError(c, "非法请求")
			c.Abort()
		}
		c.Next()
	}
}
