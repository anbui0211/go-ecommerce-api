package routers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() (r *gin.Engine) {
	r = gin.Default()
	return
}

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
