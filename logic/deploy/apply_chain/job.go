package apply_chain

import (
	"k8s-deploy/pkg/kubectl"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	v1 "k8s.io/client-go/applyconfigurations/batch/v1"
)

type Job struct {
	ctx       *ChainContext
	localYaml *v1.JobApplyConfiguration
}

func (d *Job) next(ctx *ChainContext) error {
	if !(*ctx.CdrType.APIVersion == "batch/v1" && *ctx.CdrType.Kind == "Job") {
		return nil
	}
	d.ctx = ctx

	err := d.parse()
	if err != nil {
		return err
	}
	return d.apply()
}

func (d *Job) parse() error {
	var applyYaml v1.JobApplyConfiguration
	err := yaml.Unmarshal(d.ctx.YamlByte, &applyYaml)
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
		return err
	}
	d.localYaml = &applyYaml
	return nil
}

func (d *Job) apply() error {
	_, err := kubectl.K8sClient.
		BatchV1().
		Jobs(*d.localYaml.Namespace).
		Apply(d.ctx.Ctx, d.localYaml, metav1.ApplyOptions{FieldManager: "application/apply-patch"})
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}