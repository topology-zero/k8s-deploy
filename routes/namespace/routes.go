// Code generated by goctl. DO NOT EDIT.
package namespace

import (
	"k8s-deploy/handler/namespace"
	"k8s-deploy/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterNamespaceRoute(e *gin.Engine) {
	g := e.Group("")
	g.Use(middleware.JwtMiddleware, middleware.AuthMiddleware)
	g.GET("/namespace", namespace.ListHandle)
	g.POST("/namespace", namespace.AddHandle)
	g.DELETE("/namespace", namespace.DelHandle)
}
