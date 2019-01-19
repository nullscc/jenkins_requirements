package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nullscc/jenkins_requirements/utils"
)

func Test(c *gin.Context) {
	utils.JsonRetEmpty(c)
}
