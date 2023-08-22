// Package auth 用户注册
package auth

import (
	"easy-gin/app/models/user"
	"fmt"
	"net/http"

	apiBase "easy-gin/app/controllers/api"

	"github.com/gin-gonic/gin"
)

// RegisterController 注册控制器
type RegisterController struct {
	apiBase.BaseApiController
}

// CheckPhoneExist 检测用户手机号是否已注册
func (rc *RegisterController) CheckPhoneExist(c *gin.Context) {
	//请求对象
	type CheckPhoneExistRequest struct {
		Phone string `json:"phone"`
	}

	request := CheckPhoneExistRequest{}

	//解析 请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		//打印错误信息
		fmt.Println(err.Error())
		return
	}

	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.CheckPhoneExist(request.Phone),
	})

}
