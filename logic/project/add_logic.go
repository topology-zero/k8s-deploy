package project

import (
	"k8s-deploy/model"
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
)

// Add 添加项目
func Add(ctx *svc.ServiceContext, req *types.ProjectAddRequest) error {
	projectModel := query.ProjectModel

	err := projectModel.WithContext(ctx).Create(&model.ProjectModel{
		Name:     req.Name,
		Desc:     req.Desc,
		Git:      req.Git,
		UserName: req.UserName,
		Token:    req.Token,
		UseTag:   req.UseTag,
	})

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
