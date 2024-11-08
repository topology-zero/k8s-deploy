package project

import (
	"encoding/json"
	"fmt"

	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/jinzhu/copier"
	"gorm.io/gen"
)

// List 项目列表
func List(ctx *svc.ServiceContext, req *types.ProjectListRequest) (resp types.ProjectListResponse, err error) {
	projectModel := query.ProjectModel

	var where []gen.Condition

	if len(req.Name) > 0 {
		where = append(where, projectModel.WithContext(ctx).
			Where(projectModel.Name.Like(fmt.Sprintf("%%%s%%", req.Name))).
			Or(projectModel.Desc.Like(fmt.Sprintf("%%%s%%", req.Name))))
	}

	result, count, _ := projectModel.WithContext(ctx).
		Where(where...).
		Order(projectModel.ID.Desc()).
		FindByPage((req.Page-1)*req.PageSize, req.PageSize)

	copier.Copy(&resp.Data, &result)
	for i := range result {
		var params []types.NameAndValue
		_ = json.Unmarshal([]byte(result[i].Params), &params)
		resp.Data[i].Params = params

		params = []types.NameAndValue{}
		_ = json.Unmarshal([]byte(result[i].Template), &params)
		resp.Data[i].Template = params
	}

	resp.Total = int(count)
	return
}
