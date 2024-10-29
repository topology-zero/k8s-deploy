package template

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/template"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加模板
func AddHandle(c *gin.Context) {
	var req types.TemplateAddRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := template.Add(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
