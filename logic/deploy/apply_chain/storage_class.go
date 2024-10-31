package apply_chain

import (
	"k8s-deploy/pkg/kubectl"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	v1 "k8s.io/client-go/applyconfigurations/storage/v1"
)

type StorageClass struct {
	ctx       *ChainContext
	localYaml *v1.StorageClassApplyConfiguration
}

func (d *StorageClass) next(ctx *ChainContext) error {
	if !(*ctx.CdrType.APIVersion == "storage.k8s.io/v1" && *ctx.CdrType.Kind == "StorageClass") {
		return nil
	}
	d.ctx = ctx

	err := d.parse()
	if err != nil {
		return err
	}
	return d.apply()
}

func (d *StorageClass) parse() error {
	var applyYaml v1.StorageClassApplyConfiguration
	err := yaml.Unmarshal(d.ctx.YamlByte, &applyYaml)
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
		return err
	}
	d.localYaml = &applyYaml
	return nil
}

func (d *StorageClass) apply() error {
	_, err := kubectl.K8sClient.
		StorageV1().
		StorageClasses().
		Apply(d.ctx.Ctx, d.localYaml, metav1.ApplyOptions{FieldManager: "application/apply-patch"})
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}