package template

import (
	"k8s-deploy/model"
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
)

// Add 添加模板
func Add(ctx *svc.ServiceContext, req *types.TemplateAddRequest) error {
	k8sTemplateModel := query.K8sTemplateModel

	err := k8sTemplateModel.WithContext(ctx).Create(&model.K8sTemplateModel{
		Name:    req.Name,
		Desc:    req.Desc,
		Content: req.Content,
	})

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}

	return err
}
