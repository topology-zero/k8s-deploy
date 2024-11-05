package data_chain

import (
	"encoding/json"
	"time"

	"k8s-deploy/pkg/socket"
	"k8s-deploy/query"
)

type InitMessage struct {
	ctx *ChainContext
}

func (m *InitMessage) next(ctx *ChainContext) error {
	m.ctx = ctx

	if ctx.message.MsgType != "init" {
		return nil
	}

	var id int
	_ = json.Unmarshal(ctx.message.Data, &id)

	deployLogModel := query.DeployLogModel
	data, _ := deployLogModel.WithContext(ctx.Ctx).Where(deployLogModel.Pid.Eq(id)).Order(deployLogModel.ID.Asc()).Find()
	if len(data) == 0 {
		return nil
	}

	var resp []WebsocketResponse

	for _, v := range data {
		resp = append(resp, WebsocketResponse{
			ID:         v.Pid,
			Type:       v.Type,
			Message:    v.Message,
			CreateTime: v.CreateTime.Format(time.DateTime),
		})
	}

	socket.SendJsonMessage(resp)
	return nil
}
