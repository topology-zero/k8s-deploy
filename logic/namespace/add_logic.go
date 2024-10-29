package namespace

import (
	"k8s-deploy/pkg/kubectl"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Add 添加命名空间
func Add(ctx *svc.ServiceContext, req *types.NamespaceAddRequest) error {
	ns := &v1.Namespace{}
	ns.Name = req.Name
	_, err := kubectl.K8sClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
