package common

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/common"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// ProjectParamsHandle 项目参数
func ProjectParamsHandle(c *gin.Context) {
	var req types.CommonProjectParamsRequest

	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := common.ProjectParams(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
