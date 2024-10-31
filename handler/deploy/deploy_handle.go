package deploy

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/deploy"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// DeployHandle 上线
func DeployHandle(c *gin.Context) {
	var req types.PathID
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := deploy.Deploy(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
