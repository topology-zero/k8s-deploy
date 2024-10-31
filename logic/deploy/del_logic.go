package deploy

import (
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
)

// Del 删除部署
func Del(ctx *svc.ServiceContext, req *types.PathID) error {
	deployModel := query.DeployModel

	_, err := deployModel.WithContext(ctx).Where(deployModel.ID.Eq(req.ID)).Delete()
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
