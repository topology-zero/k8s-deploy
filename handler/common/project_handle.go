package common

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/common"
	"k8s-deploy/svc"

	"github.com/gin-gonic/gin"
)

// ProjectHandle 项目列表
func ProjectHandle(c *gin.Context) {
	resp, err := common.Project(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
