package apply_chain

import (
	deploylog "k8s-deploy/logic/deploy_log"
	"k8s-deploy/pkg/kubectl"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	v1 "k8s.io/client-go/applyconfigurations/apps/v1"
)

type StatefulSet struct {
	ctx       *ChainContext
	localYaml *v1.StatefulSetApplyConfiguration
}

func (d *StatefulSet) next(ctx *ChainContext) error {
	if !(*ctx.CdrType.APIVersion == "apps/v1" && *ctx.CdrType.Kind == "StatefulSet") {
		return nil
	}
	d.ctx = ctx
	if err := d.parse(); err != nil {
		return err
	}
	if err := d.applyWarp(); err != nil {
		return err
	}
	return checkAllRunning(ctx, *d.localYaml.Namespace, d.localYaml.Spec.Template.Labels)
}

func (d *StatefulSet) parse() error {
	var applyYaml v1.StatefulSetApplyConfiguration
	err := yaml.Unmarshal(d.ctx.YamlByte, &applyYaml)
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
		return err
	}
	d.localYaml = &applyYaml
	return nil
}

func (d *StatefulSet) applyWarp() error {
	deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 0, "部署 StatefulSet \n"+string(d.ctx.YamlByte))
	err := d.apply()
	if err != nil {
		deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 3, "部署 StatefulSet 失败, 失败原因："+err.Error())
		return err
	}
	deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 1, "部署 StatefulSet 成功")
	return nil
}

func (d *StatefulSet) apply() error {
	_, err := kubectl.K8sClient.
		AppsV1().
		StatefulSets(*d.localYaml.Namespace).
		Apply(d.ctx.Ctx, d.localYaml, metav1.ApplyOptions{FieldManager: "application/apply-patch"})
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
