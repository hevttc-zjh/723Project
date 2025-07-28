package handler

import (
	"log"
	"strconv"

	"risk-insight-system/internal/service"
	"risk-insight-system/internal/utils"

	"github.com/gin-gonic/gin"
)

// GetCaseList  获取案情列表
func GetCaseList(c *gin.Context) {
	// 获取参数
	paramsStr := c.Query("params")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	log.Printf("[GetCaseList] 接收到查询请求 - 参数: %s, 分页: %s, 每页大小: %s", paramsStr, page, pageSize)
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
	caseService := service.NewCaseService()
	caseList, total, err := caseService.GetCaseList(paramsStr, pageInt, pageSizeInt)
	if err != nil {
		utils.BadRequest(c, "获取案件列表失败: "+err.Error())
		return
	}

	// 返回结果
	utils.Success(c, gin.H{
		"case_list": caseList,
		"total":     total,
		"page":      pageInt,
		"pageSize":  pageSizeInt,
	})

}

// GetCaseDetail 根据案件编号和身份证号确定唯一案件
func GetCaseDetail(c *gin.Context) {
	// 获取参数
	caseNum := c.Param("caseNum")
	idCard := c.Query("idCard")
	log.Printf("[GetCaseDetail] 接收到查询请求 - 案件编号: %s, 身份证号: %s", caseNum, idCard)
	
	// 参数校验
	if caseNum == "" {
		utils.BadRequest(c, "案件编号不能为空")
		return
	}
	
	if idCard != "" && !utils.ValidateIDCard(idCard) {
		utils.BadRequest(c, "身份证号格式错误")
		return
	}
	
	// 调用服务层方法
	caseService := service.NewCaseService()
	caseDetail, err := caseService.GetCaseDetail(caseNum, idCard)
	if err != nil {
		utils.BadRequest(c, "获取案件详情失败: " + err.Error())
		return
	}
	
	if caseDetail == nil {
		utils.NotFound(c, "未找到匹配的案件信息")
		return
	}
	
	// 返回结果
	utils.Success(c, gin.H{
		"case_detail": caseDetail,
	})
}
