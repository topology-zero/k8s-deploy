package common

import (
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"
)

// Project 项目列表
func Project(ctx *svc.ServiceContext) (resp []types.CommonProjectListResponse, err error) {
	projectModel := query.ProjectModel

	projectModel.WithContext(ctx).Select(
		projectModel.ID,
		projectModel.Name,
		projectModel.Desc,
	).Scan(&resp)

	return
}
