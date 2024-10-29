package template

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/template"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// EditHandle 编辑模板
func EditHandle(c *gin.Context) {
	var req types.TemplateEditRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := template.Edit(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
