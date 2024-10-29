package common

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/common"
	"k8s-deploy/svc"

	"github.com/gin-gonic/gin"
)

// UploadImageHandle 上传图片
func UploadImageHandle(c *gin.Context) {
	resp, err := common.UploadImage(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
