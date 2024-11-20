package apply_chain

import (
	deploylog "k8s-deploy/logic/deploy_log"
	"k8s-deploy/pkg/kubectl"

	"github.com/pkg/errors"
	networkingv1beta1 "istio.io/client-go/pkg/applyconfiguration/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

type IstioVirtualService struct {
	ctx       *ChainContext
	localYaml *networkingv1beta1.VirtualServiceApplyConfiguration
}

func (d *IstioVirtualService) next(ctx *ChainContext) error {
	if *ctx.CdrType.Kind != "VirtualService" {
		return nil
	}
	d.ctx = ctx

	err := d.parse()
	if err != nil {
		return err
	}
	return d.applyWarp()
}

func (d *IstioVirtualService) parse() error {
	var applyYaml networkingv1beta1.VirtualServiceApplyConfiguration
	err := yaml.Unmarshal(d.ctx.YamlByte, &applyYaml)
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
		return err
	}
	d.localYaml = &applyYaml
	return nil
}

func (d *IstioVirtualService) applyWarp() error {
	deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 0, "部署 IstioVirtualService \n"+string(d.ctx.YamlByte))
	err := d.apply()
	if err != nil {
		deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 3, "部署 IstioVirtualService 失败, 失败原因："+err.Error())
		return err
	}
	deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 1, "部署 IstioVirtualService 成功")
	return nil
}

func (d *IstioVirtualService) apply() error {
	_, err := kubectl.IstioClient.
		NetworkingV1beta1().
		VirtualServices(*d.localYaml.Namespace).
		Apply(d.ctx.Ctx, d.localYaml, metav1.ApplyOptions{FieldManager: "application/apply-patch"})
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
