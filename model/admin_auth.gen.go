// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAdminAuthModel = "admin_auth"

// AdminAuthModel 权限
type AdminAuthModel struct {
	ID         int            `gorm:"column:id;primaryKey;autoIncrement:true"`
	Pid        int            `gorm:"column:pid;comment:上级ID"`               // 上级ID
	Name       string         `gorm:"column:name;comment:节点名"`               // 节点名
	Key        string         `gorm:"column:key;comment:权限标识"`               // 权限标识
	IsMenu     int            `gorm:"column:is_menu;comment:是否是菜单栏 0：否 1：是"` // 是否是菜单栏 0：否 1：是
	API        string         `gorm:"column:api;comment:接口"`                 // 接口
	Action     string         `gorm:"column:action;comment:操作方法"`            // 操作方法
	CreateTime time.Time      `gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time      `gorm:"column:update_time;default:CURRENT_TIMESTAMP"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time"`
}

// TableName AdminAuthModel's table name
func (*AdminAuthModel) TableName() string {
	return TableNameAdminAuthModel
}
