package template

import (
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
)

// Edit 编辑模板
func Edit(ctx *svc.ServiceContext, req *types.TemplateEditRequest) error {
	k8sTemplateModel := query.K8sTemplateModel

	_, err := k8sTemplateModel.WithContext(ctx).
		Where(k8sTemplateModel.ID.Eq(req.ID)).
		UpdateColumnSimple(
			k8sTemplateModel.Name.Value(req.Name),
			k8sTemplateModel.Desc.Value(req.Desc),
			k8sTemplateModel.Content.Value(req.Content),
		)

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}

	return err
}
