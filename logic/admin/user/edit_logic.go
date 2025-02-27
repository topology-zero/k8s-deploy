package user

import (
	"k8s-deploy/config"
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gen/field"
)

// Edit 编辑用户
func Edit(ctx *svc.ServiceContext, req *types.UserEditRequest) error {
	if req.ID == config.SuperAdminID && req.RoleID != config.SuperAdminRoleID {
		return errors.New("无法设置超级管理员的角色")
	}

	userModel := query.AdminUserModel
	userInfo, _ := userModel.WithContext(ctx).Where(userModel.Username.Eq(req.Username), userModel.ID.Neq(req.ID)).First()
	if userInfo != nil {
		return errors.New("该账号已被注册")
	}

	userInfo, _ = userModel.WithContext(ctx).Where(userModel.Phone.Eq(req.Phone), userModel.ID.Neq(req.ID)).First()
	if userInfo != nil {
		return errors.New("该手机已被注册")
	}

	updates := []field.AssignExpr{
		userModel.Username.Value(req.Username),
		userModel.Realname.Value(req.Realname),
		userModel.Phone.Value(req.Phone),
		userModel.RoleID.Value(req.RoleID),
		userModel.Status.Value(req.Status),
	}

	if len(req.Password) > 0 {
		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.Log.Errorf("%+v", err)
			return errors.New("系统错误")
		}
		updates = append(updates, userModel.Password.Value(string(password)))
	}

	_, err := userModel.WithContext(ctx).Where(userModel.ID.Eq(req.ID)).UpdateColumnSimple(updates...)
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
