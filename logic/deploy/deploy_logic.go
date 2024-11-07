package deploy

import (
	"bufio"
	"io"
	"strings"
	"sync"

	"k8s-deploy/logic/deploy/apply_chain"
	deploylog "k8s-deploy/logic/deploy_log"
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// 一次只允许一个
var deployLock = make(map[int]*sync.Mutex)

// Deploy 上线
func Deploy(ctx *svc.ServiceContext, req *types.PathID) (err error) {
	lock, ok := deployLock[req.ID]
	if !ok {
		lock = &sync.Mutex{}
		deployLock[req.ID] = lock
	}

	if !lock.TryLock() {
		return errors.New("上线中，请不要重复上线")
	}
	defer lock.Unlock()

	deploylog.RecordStatus(ctx, req.ID, 1)

	defer func() {
		if err != nil {
			deploylog.RecordStatus(ctx, req.ID, 3)
		} else {
			deploylog.RecordStatus(ctx, req.ID, 2)
		}
	}()

	deployModel := query.DeployModel

	deployInfo, _ := deployModel.WithContext(ctx).Where(deployModel.ID.Eq(req.ID)).First()

	reader := yaml.NewYAMLReader(bufio.NewReader(strings.NewReader(deployInfo.TemplateParse)))

	for {
		rawBytes, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			deploylog.RecordLog(ctx, req.ID, 3, err.Error())
			ctx.Log.Errorf("%+v", errors.WithStack(err))
			return err
		}
		var baseType v1.TypeMetaApplyConfiguration
		err = yaml.Unmarshal(rawBytes, &baseType)
		if err != nil {
			deploylog.RecordLog(ctx, req.ID, 3, err.Error())
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
			return err
		}
	}

	return
}
