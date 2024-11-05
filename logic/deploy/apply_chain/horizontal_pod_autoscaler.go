package apply_chain

import (
	deploylog "k8s-deploy/logic/deploy_log"
	"k8s-deploy/pkg/kubectl"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	v2 "k8s.io/client-go/applyconfigurations/autoscaling/v2"
)

type HorizontalPodAutoscaler struct {
	ctx       *ChainContext
	localYaml *v2.HorizontalPodAutoscalerApplyConfiguration
}

func (d *HorizontalPodAutoscaler) next(ctx *ChainContext) error {
	if !(*ctx.CdrType.APIVersion == "autoscaling/v2" && *ctx.CdrType.Kind == "HorizontalPodAutoscaler") {
		return nil
	}
	d.ctx = ctx

	err := d.parse()
	if err != nil {
		return err
	}
	return d.applyWarp()
}

func (d *HorizontalPodAutoscaler) parse() error {
	var applyYaml v2.HorizontalPodAutoscalerApplyConfiguration
	err := yaml.Unmarshal(d.ctx.YamlByte, &applyYaml)
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
		return err
	}
	d.localYaml = &applyYaml
	return nil
}

func (d *HorizontalPodAutoscaler) applyWarp() error {
	deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 0, "部署 HorizontalPodAutoscaler \n"+string(d.ctx.YamlByte))
	err := d.apply()
	if err != nil {
		deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 3, "部署 HorizontalPodAutoscaler 失败, 失败原因："+err.Error())
		return err
	}
	deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 1, "部署 HorizontalPodAutoscaler 成功")
	return nil
}

func (d *HorizontalPodAutoscaler) apply() error {
	_, err := kubectl.K8sClient.
		AutoscalingV2().
		HorizontalPodAutoscalers(*d.localYaml.Namespace).
		Apply(d.ctx.Ctx, d.localYaml, metav1.ApplyOptions{FieldManager: "application/apply-patch"})
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
