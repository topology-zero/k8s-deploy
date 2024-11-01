package apply_chain

import (
	"time"

	"k8s-deploy/pkg/kubectl"
	"k8s-deploy/svc"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

type ChainContext struct {
	Ctx      *svc.ServiceContext
	CdrType  *v1.TypeMetaApplyConfiguration
	YamlByte []byte
}

type interactChain interface {
	// 处理逻辑
	next(ctx *ChainContext) error
}

func ApplyCdr(ctx *ChainContext) error {
	chains := []interactChain{
		// pod
		new(Deployment),
		new(ReplicationController),
		new(ReplicaSet),
		new(StatefulSet),
		new(DaemonSet),
		new(Pod),

		// job
		new(Job),
		new(CronJob),

		// service
		new(Service),
		new(Endpoint),

		// pvc
		//new(StorageClass),
		//new(PersistentVolume),
		new(PersistentVolumeClaim),

		new(ConfigMap),

		// istio
		new(IstioGateway),
		new(IstioVirtualService),
		new(IstioDestinationRule),
	}

	for _, chain := range chains {
		err := chain.next(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkAllRunning(ctx *svc.ServiceContext, namespace string, label map[string]string) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	deadline := time.After(2 * time.Minute) // 放在 for 循环外面 防止内存泄漏
	for {
		select {
		case <-deadline:
			return errors.New("POD状态错误")
		case <-ticker.C:

			selector := labels.NewSelector()
			for k, v := range label {
				match, _ := labels.NewRequirement(k, selection.Equals, []string{v})
				selector.Add(*match)
			}

			pods, err := kubectl.K8sClient.
				CoreV1().
				Pods(namespace).
				List(ctx, metav1.ListOptions{LabelSelector: selector.String()})
			if err != nil {
				ctx.Log.Errorf("%+v", errors.WithStack(err))
				return err
			}

			allRunning := true
			for _, v := range pods.Items {
				if v.Status.Phase != corev1.PodRunning {
					allRunning = false
					break
				}
			}

			if allRunning {
				return nil
			}
		}
	}
}
