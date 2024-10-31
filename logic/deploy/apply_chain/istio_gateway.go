package apply_chain

import (
	"k8s-deploy/pkg/kubectl"

	"github.com/pkg/errors"
	networkingv1alpha3 "istio.io/client-go/pkg/applyconfiguration/networking/v1alpha3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

type IstioGateway struct {
	ctx       *ChainContext
	localYaml *networkingv1alpha3.GatewayApplyConfiguration
}

func (d *IstioGateway) next(ctx *ChainContext) error {
	if !(*ctx.CdrType.APIVersion == "networking.istio.io/v1alpha3" && *ctx.CdrType.Kind == "Gateway") {
		return nil
	}
	d.ctx = ctx

	err := d.parse()
	if err != nil {
		return err
	}
	return d.apply()
}

func (d *IstioGateway) parse() error {
	var applyYaml networkingv1alpha3.GatewayApplyConfiguration
	err := yaml.Unmarshal(d.ctx.YamlByte, &applyYaml)
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
		return err
	}
	d.localYaml = &applyYaml
	return nil
}

func (d *IstioGateway) apply() error {
	_, err := kubectl.IstioClient.
		NetworkingV1alpha3().
		Gateways(*d.localYaml.Namespace).
		Apply(d.ctx.Ctx, d.localYaml, metav1.ApplyOptions{FieldManager: "application/apply-patch"})
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
