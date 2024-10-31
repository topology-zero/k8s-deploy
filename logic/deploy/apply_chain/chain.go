package apply_chain

import (
	"k8s-deploy/svc"

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
