// Package user 用户 Model
package user

import (
	"easy-gin/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel //继承 Model 基类

	Name     string `gorm:"column:name;not null;default:'';index;comment:用户名称" json:"name,omitempty"`
	Email    string `gorm:"column:email;not null;default:'';index;comment:用户邮箱" json:"-"`
	Phone    string `gorm:"column:phone;not null;default:'';index;comment:用户手机号" json:"-"`
	Password string `gorm:"column:password;not null;default:'';comment:用户密码" json:"-"`

	models.AutoTimestamp
}
