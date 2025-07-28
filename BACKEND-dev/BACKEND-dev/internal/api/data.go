package handler

import (
	"risk-insight-system/internal/utils"

	"github.com/gin-gonic/gin"
)

// ReceivePoliceData 接收警情数据
func ReceivePoliceData(c *gin.Context) {
	utils.Success(c, gin.H{
		"source": "police",
		"status": "received",
	})
}

// ReceiveSocialData 接收社会数据
func ReceiveSocialData(c *gin.Context) {
	utils.Success(c, gin.H{
		"source": "social",
		"status": "received",
	})
}

// ReceiveCaseData 接收案件数据
func ReceiveCaseData(c *gin.Context) {
	utils.Success(c, gin.H{
		"source": "case",
		"status": "received",
	})
}

// ReceiveInternalData 接收公安内部数据
func ReceiveInternalData(c *gin.Context) {
	utils.Success(c, gin.H{
		"source": "internal",
		"status": "received",
	})
}
