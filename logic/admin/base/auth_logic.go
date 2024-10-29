package base

import (
	"k8s-deploy/model"
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"
)

// Auth 获取所有权限
func Auth(ctx *svc.ServiceContext) (resp []types.BaseAuthResponse, err error) {
	authModel := query.AdminAuthModel
	auths, _ := authModel.WithContext(ctx).Find()
	return authTree(0, auths), nil
}

func authTree(pid int, auths []*model.AdminAuthModel) []types.BaseAuthResponse {
	var data []types.BaseAuthResponse
	for _, v := range auths {
		if v.Pid != pid {
			continue
		}
		data = append(data, types.BaseAuthResponse{
			ID:       v.ID,
			Pid:      v.Pid,
			Name:     v.Name,
			Children: authTree(v.ID, auths),
		})
	}
	return data
}
