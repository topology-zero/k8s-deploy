package base

import (
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"
)

// Role 获取所有角色
func Role(ctx *svc.ServiceContext) (resp []types.IDAndName, err error) {
	roleModel := query.AdminRoleModel
	_ = roleModel.WithContext(ctx).Select(roleModel.ID, roleModel.Name).Scan(&resp)
	return
}
