package base

import (
	"k8s-deploy/pkg/jwt"
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// ChangeSelfPwd 修改自己的密码
func ChangeSelfPwd(ctx *svc.ServiceContext, req *types.ChangeSelfPwdRequest) error {
	user, _ := ctx.GinContext.Get("userInfo")
	claims := user.(*jwt.Claims)

	userModel := query.AdminUserModel
	userInfo, _ := userModel.WithContext(ctx).Where(userModel.ID.Eq(claims.Id)).First()
	err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(req.OldPassword))
	if err != nil {
		return errors.New("输入的原密码错误")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		ctx.Log.Errorf("%+v", err)
		return errors.New("系统错误")
	}

	_, err = userModel.WithContext(ctx).Where(userModel.ID.Eq(claims.Id)).Update(userModel.Password, string(password))
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
