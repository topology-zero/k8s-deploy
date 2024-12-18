package user

import (
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/jinzhu/copier"
)

// List 用户列表
func List(ctx *svc.ServiceContext, req *types.UserListRequest) (resp types.UserListResponse, err error) {
	userModel := query.AdminUserModel
	roleModel := query.AdminRoleModel
	var data []types.UserList

	count, err := userModel.WithContext(ctx).LeftJoin(roleModel, userModel.RoleID.EqCol(roleModel.ID)).
		Select(userModel.ALL, roleModel.Name.As("rolename")).
		ScanByPage(&data, (req.Page-1)*req.PageSize, req.PageSize)

	resp.Total = int(count)
	copier.Copy(&resp.Data, &data)
	return
}
