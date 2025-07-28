package handler

import (
	"fmt"

	"risk-insight-system/internal/service"
	"risk-insight-system/internal/utils"

	"github.com/gin-gonic/gin"
)

// LoginHandler 处理警员登录
func LoginHandler(c *gin.Context) {
	type LoginRequest struct {
		PoliceId string `json:"policeId" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 创建用户服务
	userService := service.NewUserService()

	// 验证登录
	user, err := userService.Login(req.PoliceId, req.Password)
	if err != nil {
		utils.Unauthorized(c, "警员ID或密码错误")
		return
	}

	// 生成JWT token
	token, err := utils.GenerateJWT(
		fmt.Sprintf("%d", user.ID),
		user.PoliceId,
		user.Role,
	)
	if err != nil {
		utils.InternalServerError(c, "Token生成失敗")
		return
	}

	utils.Success(c, gin.H{
		"token": token,
		"user": gin.H{
			"id":        user.ID,
			"police_id": user.PoliceId,
			"role":      user.Role,
			"nick_name": user.NickName,
			"apartment": user.Apartment,
			"phone":     user.Phone,
			"id_card":   user.IdCard,
		},
	})
}
