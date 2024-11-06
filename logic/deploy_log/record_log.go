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
		MsgType: "log_change",
		Data: []data_chain.LogChange{{
			ID:         pid,
			Type:       tp,
			Message:    msg,
			CreateTime: time.Now().Format(time.DateTime),
		}},
	})
}

func RecordStatus(ctx *svc.ServiceContext, pid, status int) {
	deployModel := query.DeployModel
	_, err := deployModel.WithContext(ctx).
		Where(deployModel.ID.Eq(pid)).
		UpdateColumnSimple(deployModel.Status.Value(status))
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}

	socket.SendJsonMessage(data_chain.WebsocketResponse{
		MsgType: "status_change",
		Data: data_chain.StatusChange{
			PID:    pid,
			Status: status,
		},
	})
}
