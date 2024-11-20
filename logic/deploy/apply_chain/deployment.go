package apply_chain

import (
	"fmt"
	"time"

	"k8s-deploy/config"
	deploylog "k8s-deploy/logic/deploy_log"
	"k8s-deploy/pkg/kubectl"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	v1 "k8s.io/client-go/applyconfigurations/apps/v1"
)

type Deployment struct {
	ctx       *ChainContext
	localYaml *v1.DeploymentApplyConfiguration
}

func (d *Deployment) next(ctx *ChainContext) error {
	if *ctx.CdrType.Kind != "Deployment" {
		return nil
	}
	d.ctx = ctx
	if err := d.parse(); err != nil {
		return err
	}
	if err := d.applyWarp(); err != nil {
		return err
	}
	return d.checkAllRunning()
}

func (d *Deployment) parse() error {
	var applyYaml v1.DeploymentApplyConfiguration
	err := yaml.Unmarshal(d.ctx.YamlByte, &applyYaml)
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
		return err
	}
	d.localYaml = &applyYaml
	return nil
}

func (d *Deployment) applyWarp() error {
	deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 0, "部署 Deployment \n"+string(d.ctx.YamlByte))
	err := d.apply()
	if err != nil {
		deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 3, "部署 Deployment 失败, 失败原因："+err.Error())
		return err
	}
	deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 1, "部署 Deployment 成功")
	return nil
}

func (d *Deployment) apply() error {
	_, err := kubectl.K8sClient.
		AppsV1().
		Deployments(*d.localYaml.Namespace).
		Apply(d.ctx.Ctx, d.localYaml, metav1.ApplyOptions{FieldManager: "application/apply-patch"})
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}

func (d *Deployment) checkAllRunning() error {
	if config.K8sConf.WaitPod == 0 {
		return nil
	}

	deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 0, "等待 Deployment 变为 Available")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	deadline := time.After(time.Duration(config.K8sConf.WaitPod) * time.Second) // 放在 for 循环外面 防止内存泄漏
	for {
		select {
		case <-deadline:
			deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 3, "等待 Deployment 状态变化超时")
			return errors.New("Deployment 状态错误")
		case <-ticker.C:
			dp, err := kubectl.K8sClient.
				AppsV1().
				Deployments(*d.localYaml.Namespace).
				Get(d.ctx.Ctx, *d.localYaml.Name, metav1.GetOptions{})
			if err != nil {
				deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 2, "查询 Deployment 状态变化错误")
				d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
				return err
			}

			if dp.Status.AvailableReplicas-dp.Status.Replicas == 0 {
				deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 1, "Deployment 状态为 Available")
				return nil
			}

			deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 2, fmt.Sprintf("Deployment Available 状态 %d / %d", dp.Status.AvailableReplicas, dp.Status.Replicas))
		}
	}
}
