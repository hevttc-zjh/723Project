package handler

import (
	"risk-insight-system/internal/utils"

	"github.com/gin-gonic/gin"
)

// CreateFeedback 创建反馈信息
func CreateFeedback(c *gin.Context) {
	utils.Success(c, gin.H{
		"feedback_id": "feedback_001",
		"status":      "created",
	})
}

// GetPersonFeedbacks 获取人员反馈信息
func GetPersonFeedbacks(c *gin.Context) {
	personID := c.Param("person_id")
	utils.Success(c, gin.H{
		"person_id": personID,
		"feedbacks": []gin.H{},
	})
}

// UpdateFeedback 更新反馈信息
func UpdateFeedback(c *gin.Context) {
	feedbackID := c.Param("id")
	utils.Success(c, gin.H{
		"feedback_id": feedbackID,
		"status":      "updated",
	})
}
