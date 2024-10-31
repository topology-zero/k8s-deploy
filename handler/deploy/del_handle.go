package deploy

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/deploy"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// DelHandle 删除部署
func DelHandle(c *gin.Context) {
	var req types.PathID
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := deploy.Del(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
