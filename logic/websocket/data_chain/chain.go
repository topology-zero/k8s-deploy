package data_chain

import (
	"encoding/json"

	"k8s-deploy/svc"
)

type ChainContext struct {
	Ctx     *svc.ServiceContext
	message *WebsocketRequest
}

type interactChain interface {
	// 处理逻辑
	next(ctx *ChainContext) error
}

func ParseData(ctx *ChainContext, data []byte) {
	var message WebsocketRequest
	_ = json.Unmarshal(data, &message)
	ctx.message = &message

	chains := []interactChain{
		new(InitMessage),
	}

	for _, chain := range chains {
		err := chain.next(ctx)
		if err != nil {
			return
		}
	}
}
