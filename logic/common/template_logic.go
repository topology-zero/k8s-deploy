package common

import (
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"
)

// Template 模板列表
func Template(ctx *svc.ServiceContext) (resp []types.IDAndName, err error) {
	k8sTemplateModel := query.K8sTemplateModel

	k8sTemplateModel.WithContext(ctx).Select(
		k8sTemplateModel.ID,
		k8sTemplateModel.Name,
	).Scan(&resp)

	return
}
