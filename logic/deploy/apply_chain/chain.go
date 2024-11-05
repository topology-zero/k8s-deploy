package apply_chain

import (
	"sync"
	"time"

	deploylog "k8s-deploy/logic/deploy_log"
	"k8s-deploy/pkg/kubectl"
	"k8s-deploy/pkg/socket"
	"k8s-deploy/query"
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
	ID       int
}

type interactChain interface {
	// 处理逻辑
	next(ctx *ChainContext) error
}

// 一次只允许一个
var deployLock = make(map[int]*sync.Mutex)

func ApplyCdr(ctx *ChainContext) error {
	lock, ok := deployLock[ctx.ID]
	if !ok {
		lock = &sync.Mutex{}
		deployLock[ctx.ID] = lock
	}

	if !lock.TryLock() {
		return errors.New("上线中，请不要重复上线")
	}
	defer lock.Unlock()

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

	deployModel := query.DeployModel
	_, err := deployModel.WithContext(ctx.Ctx).
		Where(deployModel.ID.Eq(ctx.ID)).
		UpdateColumnSimple(deployModel.Status.Value(1))
	if err != nil {
		ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
		return err
	}

	socket.SendMessage("done")
	defer socket.SendMessage("done")

	for _, chain := range chains {
		if err = chain.next(ctx); err != nil {
			if _, dbErr := deployModel.WithContext(ctx.Ctx).
				Where(deployModel.ID.Eq(ctx.ID)).
				UpdateColumnSimple(deployModel.Status.Value(3)); dbErr != nil {
				ctx.Ctx.Log.Errorf("%+v", errors.WithStack(dbErr))
			}
			return err
		}
	}

	_, err = deployModel.WithContext(ctx.Ctx).
		Where(deployModel.ID.Eq(ctx.ID)).
		UpdateColumnSimple(deployModel.Status.Value(2))
	if err != nil {
		ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}

func checkAllRunning(ctx *ChainContext, namespace string, label map[string]string) error {
	deploylog.RecordLog(ctx.Ctx, ctx.ID, 0, "等待POD变为PodRunning")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	deadline := time.After(20 * time.Second) // 放在 for 循环外面 防止内存泄漏
	for {
		select {
		case <-deadline:
			deploylog.RecordLog(ctx.Ctx, ctx.ID, 3, "等待POD状态变化超时")
			return errors.New("POD状态错误")
		case <-ticker.C:

			selector := labels.NewSelector()
			for k, v := range label {
				match, _ := labels.NewRequirement(k, selection.Equals, []string{v})
				selector = selector.Add(*match)
			}

			pods, err := kubectl.K8sClient.
				CoreV1().
				Pods(namespace).
				List(ctx.Ctx, metav1.ListOptions{LabelSelector: selector.String()})
			if err != nil {
				deploylog.RecordLog(ctx.Ctx, ctx.ID, 2, "查询POD状态变化错误")
				ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
				return err
			}

			allRunning := true
			for _, v := range pods.Items {
				if v.Status.Phase != corev1.PodRunning {
					deploylog.RecordLog(ctx.Ctx, ctx.ID, 2, "POD状态为"+string(v.Status.Phase))
					allRunning = false
					break
				}
			}

			if allRunning {
				deploylog.RecordLog(ctx.Ctx, ctx.ID, 1, "POD状态为Running")
				return nil
			}
		}
	}
}
