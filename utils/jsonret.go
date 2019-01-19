package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/nullscc/jenkins_requirements/models"
	"net/http"
)

func JsonRet(c *gin.Context, info interface{}) {
	response := models.Response{
		Code: 200,
		Info: info,
	}
	c.JSON(http.StatusOK, &response)
}

func JsonRetEmpty(c *gin.Context) {
	response := models.Response{
		Code: 200,
		Info: struct{}{},
	}
	c.JSON(http.StatusOK, &response)
}

func JsonRetTip(c *gin.Context, info interface{}, msg string) {
	response := models.Response{
		Code: 208,
		Msg:  msg,
		Info: info,
	}
	c.JSON(http.StatusOK, &response)
}

func JsonRetWarn(c *gin.Context, msg string) {
	response := models.Response{
		Code: 209,
		Msg:  msg,
		Info: struct{}{},
	}
	c.JSON(http.StatusOK, &response)
}

func JsonRetError(c *gin.Context, msg string) {
	response := models.Response{
		Code: 500,
		Msg:  msg,
		Info: struct{}{},
	}
	c.JSON(http.StatusOK, &response)
}

func JsonRetTokenExpire(c *gin.Context, msg string) {
	response := models.Response{
		Code: 20002,
		Msg:  msg,
		Info: struct{}{},
	}
	c.JSON(http.StatusOK, &response)
}
