// Code generated by goctl. DO NOT EDIT.
package common

import (
	"k8s-deploy/handler/common"
	"k8s-deploy/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCommonRoute(e *gin.Engine) {
	g := e.Group("")
	g.Use(middleware.JwtMiddleware)
	g.POST("/upload/image", common.UploadImageHandle)
}
