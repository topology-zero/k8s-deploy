package base

import (
	"k8s-deploy/pkg/jwt"
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"
)

// UserInfo 获取用户信息
func UserInfo(ctx *svc.ServiceContext) (resp types.UserInfoResponse, err error) {
	user, _ := ctx.GinContext.Get("userInfo")
	claims := user.(*jwt.Claims)
	userModel := query.AdminUserModel
	resp = userModel.WithContext(ctx).GetUserInfo(claims.Id)
	if claims.RoleId == 1 {
		resp.Authkeys = "*"
	}
	return
}
