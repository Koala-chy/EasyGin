// Package user 用户 Model
package user

import (
	"easy-gin/app/models"
	"easy-gin/pkg/database"
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

// CheckEmailExist 判断邮箱是否注册
func CheckEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// CheckPhoneExist 判断手机号是否已注册
func CheckPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
