// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDeployLogModel = "deploy_log"

// DeployLogModel 部署日志
type DeployLogModel struct {
	ID         int       `gorm:"column:id;primaryKey;autoIncrement:true"`
	Pid        int       `gorm:"column:pid;comment:部署ID"`                                 // 部署ID
	Type       int       `gorm:"column:type;comment:消息登记 0：默认 1：成功 2：警告 3：错误"`            // 消息登记 0：默认 1：成功 2：警告 3：错误
	Message    string    `gorm:"column:message;comment:消息"`                               // 消息
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:时间"` // 时间
}

// TableName DeployLogModel's table name
func (*DeployLogModel) TableName() string {
	return TableNameDeployLogModel
}
