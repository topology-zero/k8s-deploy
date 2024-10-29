package template

import (
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/jinzhu/copier"
)

// List 模板列表
func List(ctx *svc.ServiceContext, req *types.TemplateListRequest) (resp types.TemplateListResponse, err error) {
	k8sTemplateModel := query.K8sTemplateModel

	data, count, _ := k8sTemplateModel.WithContext(ctx).
		Order(k8sTemplateModel.ID.Desc()).
		FindByPage((req.Page-1)*req.PageSize, req.PageSize)

	copier.Copy(&resp.Data, &data)
	resp.Total = int(count)
	return
}
