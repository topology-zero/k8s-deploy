package project

import (
	"encoding/json"

	"k8s-deploy/model"
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
)

// Add 添加项目
func Add(ctx *svc.ServiceContext, req *types.ProjectAddRequest) error {
	projectModel := query.ProjectModel

	params, _ := json.Marshal(req.Params)
	template, _ := json.Marshal(req.Template)

	err := projectModel.WithContext(ctx).Create(&model.ProjectModel{
		Name:     req.Name,
		Desc:     req.Desc,
		Git:      req.Git,
		UserName: req.UserName,
		Token:    req.Token,
		UseTag:   req.UseTag,
		Params:   string(params),
		Template: string(template),
	})

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
