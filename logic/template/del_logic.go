package template

import (
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
)

// Del 删除模板
func Del(ctx *svc.ServiceContext, req *types.PathID) error {
	k8sTemplateModel := query.K8sTemplateModel

	_, err := k8sTemplateModel.WithContext(ctx).Where(k8sTemplateModel.ID.Eq(req.ID)).Delete()

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}

	return err
}
