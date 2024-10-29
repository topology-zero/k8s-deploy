package login

import (
	"strings"

	"k8s-deploy/config"
	"k8s-deploy/pkg/jwt"
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Login 登录
func Login(ctx *svc.ServiceContext, req *types.LoginRequest) (resp types.LoginResponse, err error) {
	req.Code = strings.ToLower(req.Code)
	match := config.Captcha.Verify(req.CodeID, req.Code, true)
	if !match {
		err = errors.New("验证码错误")
		return
	}

	userModel := query.AdminUserModel
	userInfo, _ := userModel.WithContext(ctx).Where(userModel.Username.Eq(req.Username)).First()
	if userInfo == nil {
		err = errors.New("用户名或密码错误")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(req.Password))
	if err != nil {
		err = errors.New("用户名或密码错误")
		return
	}

	if userInfo.Status != 1 {
		err = errors.New("该用户已被管理员停用")
		return
	}

	token, err := jwt.MakeToken(userInfo.ID, userInfo.RoleID, userInfo.Phone)
	if err != nil {
		err = errors.New("系统错误")
		return
	}
	resp.Jwt = token
	return
}
