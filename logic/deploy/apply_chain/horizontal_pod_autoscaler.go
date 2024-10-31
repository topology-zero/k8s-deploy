package apply_chain

import (
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
	return d.apply()
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
