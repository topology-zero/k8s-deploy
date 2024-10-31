package apply_chain

import (
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
	if !(*ctx.CdrType.APIVersion == "v1" && *ctx.CdrType.Kind == "PersistentVolumeClaim") {
		return nil
	}
	d.ctx = ctx

	err := d.parse()
	if err != nil {
		return err
	}
	return d.apply()
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

func (d *PersistentVolumeClaim) apply() error {
	_, err := kubectl.K8sClient.
		CoreV1().
		PersistentVolumeClaims(*d.localYaml.Namespace).
		Apply(d.ctx.Ctx, d.localYaml, metav1.ApplyOptions{FieldManager: "application/apply-patch"})
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
