package deploy

import (
	"encoding/json"
	"time"

	"k8s-deploy/model"
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"gorm.io/gen"
)

// List 部署列表
func List(ctx *svc.ServiceContext, req *types.DeployListRequest) (resp types.DeployListResponse, err error) {
	deployModel := query.DeployModel

	var where []gen.Condition

	if len(req.Name) > 0 {
		where = append(where, deployModel.Name.Like("%"+req.Name+"%"))
	}

	data, count, _ := deployModel.WithContext(ctx).
		Where(where...).
		Order(deployModel.ID.Desc()).
		FindByPage((req.Page-1)*req.PageSize, req.PageSize)

	resp.Total = int(count)

	for _, v := range data {
		var project model.ProjectModel
		var template model.K8sTemplateModel

		json.Unmarshal([]byte(v.Project), &project)
		json.Unmarshal([]byte(v.Template), &template)

		resp.Data = append(resp.Data, types.DeployList{
			ID:           v.ID,
			DeployName:   v.Name,
			ProjectName:  project.Name,
			TemplateName: template.Name,
			Status:       v.Status,
			UpdateTime:   v.UpdateTime.Format(time.DateTime),
		})

	}

	return
}
