package deploy_log

import (
	"time"

	"k8s-deploy/logic/websocket/data_chain"
	"k8s-deploy/model"
	"k8s-deploy/pkg/socket"
	"k8s-deploy/query"
	"k8s-deploy/svc"

	"github.com/pkg/errors"
)

func RecordLog(ctx *svc.ServiceContext, pid, tp int, msg string) {
	deployLogModel := query.DeployLogModel

	err := deployLogModel.WithContext(ctx).Create(&model.DeployLogModel{
		Pid:     pid,
		Type:    tp,
		Message: msg,
	})

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}

	socket.SendJsonMessage(data_chain.WebsocketResponse{
		ID:         pid,
		Type:       tp,
		Message:    msg,
		CreateTime: time.Now().Format(time.DateTime),
	})
}
