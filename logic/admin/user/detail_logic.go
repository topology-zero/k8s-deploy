package user

import (
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/jinzhu/copier"
)

// Detail 用户详情
func Detail(ctx *svc.ServiceContext, req *types.PathID) (resp types.UserDetailResponse, err error) {
	userModel := query.AdminUserModel
	userInfo, _ := userModel.WithContext(ctx).Where(userModel.ID.Eq(req.ID)).First()
	copier.Copy(&resp, &userInfo)
	return
}
