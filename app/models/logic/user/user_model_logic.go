// Package user 用户模型 logic 层
package user

import (
	"easy-gin/app/models/entity/user"
	"easy-gin/pkg/database"
)

// CheckEmailExist 判断邮箱是否注册
func CheckEmailExist(email string) bool {
	var count int64
	database.DB.Model(user.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// CheckPhoneExist 判断手机号是否已注册
func CheckPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(user.User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
