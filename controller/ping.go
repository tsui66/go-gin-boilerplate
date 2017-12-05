package controller

import (
	"github.com/gin-gonic/gin"
)

type PingController struct{}

func (p PingController) Status(c *gin.Context) {
	c.JSON(200, gin.H{ "message": "I'm OK!" })
	return
}