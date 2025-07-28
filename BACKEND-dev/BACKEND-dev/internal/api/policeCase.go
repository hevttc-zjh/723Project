package handler

import (
	"log"
	"risk-insight-system/internal/service"
	"risk-insight-system/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取警情列表
func GetPoliceCaseList(c *gin.Context) {
	// 获取参数
	paramsStr := c.Query("params")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	log.Printf("[GetPersonList] 接收到查询请求 - 参数: %s, 分页: %s, 每页大小: %s", paramsStr, page, pageSize)
	// 解析分页参数
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 1 || pageSizeInt > 100 {
		pageSizeInt = 10
	}
	// 调用服务层方法
	policeCaseService := service.NewPoliceCaseService()
	policeCaseList, total, err := policeCaseService.GetPoliceCaseList(paramsStr, pageInt, pageSizeInt)
	if err != nil {
		utils.BadRequest(c, "查询失败: "+err.Error())
		log.Printf("[GetPoliceCaseList] 查询失败: %v", err)
		return
	}
	
	utils.Success(c, gin.H{
		"police_case_list": policeCaseList,
		"total":            total,
		"page":				pageInt,
		"pageSize":			pageSizeInt,
	})
}

// 获取警情详情
func GetPoliceCaseDetail(c *gin.Context) {
	utils.Success(c, gin.H{
		"police_case_detail": gin.H{},
	})
}
