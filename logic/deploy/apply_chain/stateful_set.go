package apply_chain

import (
	"time"

	"k8s-deploy/pkg/kubectl"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
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

	err := d.parse()
	if err != nil {
		return err
	}
	return d.apply()
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

func (d *StatefulSet) apply() error {
	_, err := kubectl.K8sClient.
		AppsV1().
		StatefulSets(*d.localYaml.Namespace).
		Apply(d.ctx.Ctx, d.localYaml, metav1.ApplyOptions{FieldManager: "application/apply-patch"})
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
		return err
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	deadline := time.After(2 * time.Minute) // 放在 for 循环外面 防止内存泄漏
	for {
		select {
		case <-deadline:
			return errors.New("POD状态错误")
		case <-ticker.C:
			deployment, err := kubectl.K8sClient.
				CoreV1().
				Pods(*d.localYaml.Namespace).
				Get(d.ctx.Ctx, *d.localYaml.Spec.Template.Name, metav1.GetOptions{})
			if err != nil {
				d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
				return err
			}

			if deployment.Status.Phase == corev1.PodRunning {
				return nil
			}
		}
	}
}
