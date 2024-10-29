package namespace

import (
	"k8s-deploy/pkg/kubectl"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Del 删除命名空间
func Del(ctx *svc.ServiceContext, req *types.NamespaceDelRequest) error {
	err := kubectl.K8sClient.CoreV1().Namespaces().Delete(ctx, req.Name, metav1.DeleteOptions{})
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
