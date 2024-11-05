package deploy

import (
	"bufio"
	"io"
	"strings"

	"k8s-deploy/logic/deploy/apply_chain"
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// Deploy 上线
func Deploy(ctx *svc.ServiceContext, req *types.PathID) error {
	deployModel := query.DeployModel

	deployInfo, _ := deployModel.WithContext(ctx).Where(deployModel.ID.Eq(req.ID)).First()

	reader := yaml.NewYAMLReader(bufio.NewReader(strings.NewReader(deployInfo.TemplateParse)))

	for {
		rawBytes, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			ctx.Log.Errorf("%+v", errors.WithStack(err))
			return err
		}
		var baseType v1.TypeMetaApplyConfiguration
		err = yaml.Unmarshal(rawBytes, &baseType)
		if err != nil {
			ctx.Log.Errorf("%+v", errors.WithStack(err))
			return err
		}

		ChainCtx := apply_chain.ChainContext{
			Ctx:      ctx,
			CdrType:  &baseType,
			YamlByte: rawBytes,
			ID:       req.ID,
		}
		err = apply_chain.ApplyCdr(&ChainCtx)
		if err != nil {
			ctx.Log.Errorf("%+v", errors.WithStack(err))
			return err
		}
	}

	return nil
}
