package apply_chain

import (
	"k8s-deploy/pkg/kubectl"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

type PersistentVolume struct {
	ctx       *ChainContext
	localYaml *v1.PersistentVolumeApplyConfiguration
}

func (d *PersistentVolume) next(ctx *ChainContext) error {
	if !(*ctx.CdrType.APIVersion == "v1" && *ctx.CdrType.Kind == "PersistentVolume") {
		return nil
	}
	d.ctx = ctx

	err := d.parse()
	if err != nil {
		return err
	}
	return d.apply()
}

func (d *PersistentVolume) parse() error {
	var applyYaml v1.PersistentVolumeApplyConfiguration
	err := yaml.Unmarshal(d.ctx.YamlByte, &applyYaml)
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
		return err
	}
	d.localYaml = &applyYaml
	return nil
}

func (d *PersistentVolume) apply() error {
	_, err := kubectl.K8sClient.
		CoreV1().
		PersistentVolumes().
		Apply(d.ctx.Ctx, d.localYaml, metav1.ApplyOptions{FieldManager: "application/apply-patch"})
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
