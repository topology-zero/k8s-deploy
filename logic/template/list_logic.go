package template

import (
	"fmt"

	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/jinzhu/copier"
	"gorm.io/gen"
)

// List 模板列表
func List(ctx *svc.ServiceContext, req *types.TemplateListRequest) (resp types.TemplateListResponse, err error) {
	k8sTemplateModel := query.K8sTemplateModel

	var where []gen.Condition

	if len(req.Name) > 0 {
		where = append(where, k8sTemplateModel.WithContext(ctx).
			Where(k8sTemplateModel.Name.Like(fmt.Sprintf("%%%s%%", req.Name))).
			Or(k8sTemplateModel.Desc.Like(fmt.Sprintf("%%%s%%", req.Name))))
	}

	data, count, _ := k8sTemplateModel.WithContext(ctx).
		Where(where...).
		Order(k8sTemplateModel.ID.Desc()).
		FindByPage((req.Page-1)*req.PageSize, req.PageSize)

	copier.Copy(&resp.Data, &data)
	resp.Total = int(count)
	return
}
