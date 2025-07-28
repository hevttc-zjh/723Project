package handler

import (
	"risk-insight-system/internal/utils"

	"github.com/gin-gonic/gin"
)

// GetDashboardStats 获取首页大屏统计数据
func GetDashboardStats(c *gin.Context) {
	utils.Success(c, gin.H{
		"total_persons":       0,
		"high_risk_persons":   0,
		"medium_risk_persons": 0,
		"low_risk_persons":    0,
		"today_alerts":        0,
		"weekly_alerts":       0,
	})
}

// GetPersonDistribution 获取人员分布统计
func GetPersonDistribution(c *gin.Context) {
	utils.Success(c, gin.H{
		"by_district":   []gin.H{},
		"by_risk_level": []gin.H{},
	})
}

// GetRiskLevelStats 获取风险等级统计
func GetRiskLevelStats(c *gin.Context) {
	utils.Success(c, gin.H{
		"risk_levels": []gin.H{
			{"level": "high", "count": 0, "percentage": 0},
			{"level": "medium", "count": 0, "percentage": 0},
			{"level": "low", "count": 0, "percentage": 0},
		},
	})
}
