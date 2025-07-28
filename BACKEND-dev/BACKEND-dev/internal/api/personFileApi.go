package handler

import (
	"risk-insight-system/internal/service"
	"risk-insight-system/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PersonFileApi 人员档案接口
func PersonFileApi(c *gin.Context) {
	type fileRequest struct {
		IdCard string `form:"IdCard"`
		Name   string `form:"Name"`
		Phone  string `form:"Phone"`
	}
	var req fileRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.BadRequest(c, "参数错误")
		return
	}
	if req.IdCard == "" && req.Name == "" && req.Phone == "" {
		utils.BadRequest(c, "请至少提供身份证号、姓名或手机号中的一个")
		return
	}
	personFileService := service.NewPersonFileService()
	personFile, err := personFileService.GetPersonFile(req.IdCard, req.Name, req.Phone)
	if err != nil {
		utils.BadRequest(c, "查询失败: "+err.Error())
		return
	}
	if personFile == nil {
		utils.NotFound(c, "未找到人员信息")
		return
	}
	// 判断返回内容类型
	if personFile.PersonInfo != nil {
		utils.Success(c, gin.H{"person_info": personFile.PersonInfo})
	} else {
		utils.Success(c, gin.H{"phone": personFile.Phone, "department": personFile.Department})
	}
}

// GetPersonPhoneByIDCardApi 通过身份证号查手机号和部门
func GetPersonPhoneByIDCardApi(c *gin.Context) {
	idCard := c.Query("IdCard")
	if idCard == "" {
		utils.BadRequest(c, "请提供身份证号IdCard参数")
		return
	}
	personFileService := service.NewPersonFileService()
	personFile, err := personFileService.GetPersonPhoneByIDCard(idCard)
	if err != nil {
		utils.BadRequest(c, "查询失败: "+err.Error())
		return
	}
	if personFile == nil {
		utils.NotFound(c, "未找到人员信息")
		return
	}
	utils.Success(c, gin.H{"phone": personFile.Phone, "department": personFile.Department})
}

// GetPersonTagsByIDCardApi 通过身份证号查找标签
func GetPersonTagsByIDCardApi(c *gin.Context) {
	idCard := c.Query("IdCard")
	if idCard == "" {
		utils.BadRequest(c, "请提供身份证号IdCard参数")
		return
	}
	personFileService := service.NewPersonFileService()
	tags, err := personFileService.GetPersonTagsByIDCard(idCard)
	if err != nil {
		utils.BadRequest(c, "查询失败: "+err.Error())
		return
	}
	utils.Success(c, gin.H{"tags": tags})
}

// GetPoliceCaseListByPersonFileApi 调用警情列表接口
func GetPoliceCaseListByPersonFileApi(c *gin.Context) {
	paramsStr := c.Query("params")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 1 || pageSizeInt > 100 {
		pageSizeInt = 10
	}
	personFileService := service.NewPersonFileService()
	policeCaseList, total, err := personFileService.GetPoliceCaseList(paramsStr, pageInt, pageSizeInt)
	if err != nil {
		utils.BadRequest(c, "查询失败: "+err.Error())
		return
	}
	utils.Success(c, gin.H{
		"police_case_list": policeCaseList,
		"total":            total,
		"page":             pageInt,
		"pageSize":         pageSizeInt,
	})
}

// GetCaseListByPersonFileApi 调用案件列表接口
func GetCaseListByPersonFileApi(c *gin.Context) {
	paramsStr := c.Query("params")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 1 || pageSizeInt > 100 {
		pageSizeInt = 10
	}
	personFileService := service.NewPersonFileService()
	caseList, total, err := personFileService.GetCaseList(paramsStr, pageInt, pageSizeInt)
	if err != nil {
		utils.BadRequest(c, "查询失败: "+err.Error())
		return
	}
	utils.Success(c, gin.H{
		"case_list": caseList,
		"total":     total,
		"page":      pageInt,
		"pageSize":  pageSizeInt,
	})
}

// GetPersonWithMedicalByIDCardApi 通过身份证号联合查询人员与医疗信息
func GetPersonWithMedicalByIDCardApi(c *gin.Context) {
	idCard := c.Query("IdCard")
	if idCard == "" {
		utils.BadRequest(c, "请提供身份证号IdCard参数")
		return
	}
	personFileService := service.NewPersonFileService()
	result, err := personFileService.GetPersonWithMedicalByIDCard(idCard)
	if err != nil {
		utils.BadRequest(c, "查询失败: "+err.Error())
		return
	}
	if result == nil {
		utils.NotFound(c, "未找到相关信息")
		return
	}
	utils.Success(c, result)
}
