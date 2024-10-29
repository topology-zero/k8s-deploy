package project

import (
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
)

// Del 删除项目
func Del(ctx *svc.ServiceContext, req *types.PathID) error {
	projectModel := query.ProjectModel

	_, err := projectModel.WithContext(ctx).Where(projectModel.ID.Eq(req.ID)).Delete()

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
