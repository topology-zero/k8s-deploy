package login

import (
	"k8s-deploy/config"
	"k8s-deploy/svc"
	"k8s-deploy/types"
)

// Code 获取验证码
func Code(_ *svc.ServiceContext) (resp types.CodeResponse, err error) {
	resp.ID, resp.Image, err = config.Captcha.Generate()
	return
}
