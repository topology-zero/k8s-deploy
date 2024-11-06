package ui

import (
	"github.com/gin-gonic/gin"
)

func Setup(c *gin.Engine) {
	c.StaticFile("/", "./ui")
	c.Static("/static", "./ui/static")
}
