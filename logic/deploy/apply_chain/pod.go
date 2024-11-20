package apply_chain

import (
	deploylog "k8s-deploy/logic/deploy_log"
	"k8s-deploy/pkg/kubectl"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

type Pod struct {
	ctx       *ChainContext
	localYaml *v1.PodApplyConfiguration
}

func (d *Pod) next(ctx *ChainContext) error {
	if *ctx.CdrType.Kind != "Pod" {
		return nil
	}
	d.ctx = ctx
	if err := d.parse(); err != nil {
		return err
	}
	if err := d.applyWarp(); err != nil {
		return err
	}
	return checkAllRunning(ctx, *d.localYaml.Namespace, d.localYaml.Labels)
}

func (d *Pod) parse() error {
	var applyYaml v1.PodApplyConfiguration
	err := yaml.Unmarshal(d.ctx.YamlByte, &applyYaml)
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
		return err
	}
	d.localYaml = &applyYaml
	return nil
}

func (d *Pod) applyWarp() error {
	deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 0, "部署 Pod \n"+string(d.ctx.YamlByte))
	err := d.apply()
	if err != nil {
		deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 3, "部署 Pod 失败, 失败原因："+err.Error())
		return err
	}
	deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 1, "部署 Pod 成功")
	return nil
}

func (d *Pod) apply() error {
	_, err := kubectl.K8sClient.
		CoreV1().
		Pods(*d.localYaml.Namespace).
		Apply(d.ctx.Ctx, d.localYaml, metav1.ApplyOptions{FieldManager: "application/apply-patch"})
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
