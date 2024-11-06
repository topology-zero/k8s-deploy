package data_chain

import "encoding/json"

type WebsocketRequest struct {
	MsgType string          `json:"msgType"`
	Data    json.RawMessage `json:"data"`
}

type WebsocketResponse struct {
	MsgType string `json:"msgType"`
	Data    any    `json:"data"`
}

type LogChange struct {
	ID         int    `json:"id"`
	Type       int    `json:"type"`
	Message    string `json:"message"`
	CreateTime string `json:"createTime"`
}

type StatusChange struct {
	PID    int `json:"pid"`
	Status int `json:"status"`
}
