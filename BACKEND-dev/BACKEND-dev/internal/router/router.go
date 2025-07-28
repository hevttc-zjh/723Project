package router

import (
	"risk-insight-system/config"
	handler "risk-insight-system/internal/api"
	"risk-insight-system/internal/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 设置Gin模式
	gin.SetMode(config.GetString("server.mode"))

	r := gin.New()

	// 使用中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())

	//健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "情指行风险洞察系统运行正常",
		})
	})

	// API版本分组
	v1 := r.Group("/api/v1")
	{
		// 用户登录相关接口
		loginGroup := v1.Group("/auth")
		{
			loginGroup.POST("/login", handler.LoginHandler)
		}
		// 人员画像相关接口
		personGroup := v1.Group("/person")
		{
			personGroup.GET("/search", handler.SearchPerson)
			personGroup.GET("/:id", handler.GetPersonProfile)
			personGroup.GET("/:id/relations", handler.GetPersonRelations)
			personGroup.GET("/:id/cases", handler.GetPersonCases)
			personGroup.GET("/list", handler.GetPersonList)
			personGroup.GET("/file", handler.PersonFileApi) // 新增人员档案接口
			personGroup.GET("/phone-by-idcard", handler.GetPersonPhoneByIDCardApi)
			personGroup.GET("/tags-by-idcard", handler.GetPersonTagsByIDCardApi)
			personGroup.GET("/police-case-list-by-file", handler.GetPoliceCaseListByPersonFileApi)
			personGroup.GET("/case-list-by-file", handler.GetCaseListByPersonFileApi)
			personGroup.GET("/person-with-medical", handler.GetPersonWithMedicalByIDCardApi)
		}

		// 数据接入相关接口
		dataGroup := v1.Group("/data")
		{
			dataGroup.POST("/police", handler.ReceivePoliceData)
			dataGroup.POST("/social", handler.ReceiveSocialData)
			dataGroup.POST("/case", handler.ReceiveCaseData)
			dataGroup.POST("/internal", handler.ReceiveInternalData)
		}

		// 反馈信息相关接口
		feedbackGroup := v1.Group("/feedback")
		{
			feedbackGroup.POST("/", handler.CreateFeedback)
			feedbackGroup.GET("/:person_id", handler.GetPersonFeedbacks)
			feedbackGroup.PUT("/:id", handler.UpdateFeedback)
		}

		// 统计分析相关接口
		statsGroup := v1.Group("/stats")
		{
			statsGroup.GET("/dashboard", handler.GetDashboardStats)
			statsGroup.GET("/distribution", handler.GetPersonDistribution)
			statsGroup.GET("/risk-levels", handler.GetRiskLevelStats)
		}
		// 警情相关接口
		policeCaseGroup := v1.Group("/police/case")
		{
			policeCaseGroup.GET("/list", handler.GetPoliceCaseList)
			policeCaseGroup.GET("/:id", handler.GetPoliceCaseDetail)
		}
		// 案情相关接口
		caseGroup := v1.Group("/case")
		{
			caseGroup.GET("/list", handler.GetCaseList)
			caseGroup.GET("/detail", handler.GetCaseDetail)
		}
	}

	return r
}
