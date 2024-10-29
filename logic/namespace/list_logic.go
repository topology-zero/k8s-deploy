package namespace

import (
	"k8s-deploy/pkg/kubectl"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// List 命名空间列表
func List(ctx *svc.ServiceContext) (resp []types.NamespaceListResponse, err error) {
	ns, err := kubectl.K8sClient.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		return
	}
	for _, item := range ns.Items {
		resp = append(resp, types.NamespaceListResponse{
			Name: item.Name,
		})
	}
	return
}
