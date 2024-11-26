package apply_chain

import (
	deploylog "k8s-deploy/logic/deploy_log"
	"k8s-deploy/pkg/kubectl"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

type PersistentVolumeClaim struct {
	ctx       *ChainContext
	localYaml *v1.PersistentVolumeClaimApplyConfiguration
}

func (d *PersistentVolumeClaim) next(ctx *ChainContext) error {
	if *ctx.CdrType.Kind != "PersistentVolumeClaim" {
		return nil
	}
	d.ctx = ctx

	err := d.parse()
	if err != nil {
		return err
	}
	return d.applyWarp()
}

func (d *PersistentVolumeClaim) parse() error {
	var applyYaml v1.PersistentVolumeClaimApplyConfiguration
	err := yaml.Unmarshal(d.ctx.YamlByte, &applyYaml)
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
		return err
	}
	d.localYaml = &applyYaml
	return nil
}

func (d *PersistentVolumeClaim) applyWarp() error {
	deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 0, "部署 PersistentVolumeClaim \n"+string(d.ctx.YamlByte))
	err := d.apply()
	if err != nil {
		deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 3, "部署 PersistentVolumeClaim 失败, 失败原因："+err.Error())
		return err
	}
	deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 1, "部署 PersistentVolumeClaim 成功")
	return nil
}

func (d *PersistentVolumeClaim) apply() error {
	_, err := kubectl.K8sClient.
		CoreV1().
		PersistentVolumeClaims(*d.localYaml.Namespace).
		Apply(d.ctx.Ctx, d.localYaml, metav1.ApplyOptions{FieldManager: "application/apply-patch", DryRun: d.ctx.DryRun})
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
