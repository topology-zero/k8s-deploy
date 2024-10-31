package common

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/common"
	"k8s-deploy/svc"

	"github.com/gin-gonic/gin"
)

// TemplateHandle 模板列表
func TemplateHandle(c *gin.Context) {
	resp, err := common.Template(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
