package template

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/template"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// ListHandle 模板列表
func ListHandle(c *gin.Context) {
	var req types.TemplateListRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := template.List(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
