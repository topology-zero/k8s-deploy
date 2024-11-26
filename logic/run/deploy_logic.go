package run

import (
	"bufio"
	"io"
	"strings"
	"sync"

	"k8s-deploy/logic/deploy/apply_chain"
	"k8s-deploy/logic/deploy_log"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

var runLock sync.Mutex

// Deploy 运行
func Deploy(ctx *svc.ServiceContext, req *types.RunDeployRequest) (err error) {
	if !runLock.TryLock() {
		return errors.New("上线中，请不要重复上线")
	}
	defer runLock.Unlock()

	deploy_log.RecordStatus(ctx, -1, 1)

	defer func() {
		if err != nil {
			deploy_log.RecordStatus(ctx, -1, 3)
		} else {
			deploy_log.RecordStatus(ctx, -1, 2)
		}
	}()

	reader := yaml.NewYAMLReader(bufio.NewReader(strings.NewReader(req.Yaml)))

	for {
		rawBytes, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			deploy_log.RecordLog(ctx, -1, 3, err.Error())
			ctx.Log.Errorf("%+v", errors.WithStack(err))
			return err
		}
		var baseType v1.TypeMetaApplyConfiguration
		err = yaml.Unmarshal(rawBytes, &baseType)
		if err != nil {
			deploy_log.RecordLog(ctx, -1, 3, err.Error())
			ctx.Log.Errorf("%+v", errors.WithStack(err))
			return err
		}

		if baseType.Kind == nil || baseType.APIVersion == nil {
			err = errors.New("yaml 解析错误")
			deploy_log.RecordLog(ctx, -1, 3, err.Error())
			ctx.Log.Errorf("%+v", err)
			return err
		}

		ChainCtx := apply_chain.ChainContext{
			Ctx:      ctx,
			CdrType:  &baseType,
			YamlByte: rawBytes,
			ID:       -1,
		}
		if req.IsDebug {
			ChainCtx.DryRun = []string{metav1.DryRunAll}
		}

		err = apply_chain.ApplyCdr(&ChainCtx)
		if err != nil {
			return err
		}
	}
	return nil
}
