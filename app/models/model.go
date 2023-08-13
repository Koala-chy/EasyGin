// Package models 模型公共属性方法
package models

import "time"

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// AutoTimestamp 模型自动更新时间
type AutoTimestamp struct {
	CreatedAt time.Time `gorm:"column:created_at;default:null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:created_at;default:null" json:"updated_at,omitempty"`
}
