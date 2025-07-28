package handler

import (
	"fmt"
	"log"
	"strconv"

	"risk-insight-system/internal/model"
	"risk-insight-system/internal/service"
	"risk-insight-system/internal/utils"

	"github.com/gin-gonic/gin"
)

// SearchPerson 搜索人员
func SearchPerson(c *gin.Context) {
	type searchRequest struct {
		IdCard string `form:"IdCard" binding:"omitempty"`
	}
	var req searchRequest
	fmt.Println(req.IdCard)
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.BadRequest(c, "参数错误")
		log.Printf("[SearchPerson] 参数绑定失败: %v", err)
		return
	}
	log.Printf("[SearchPerson] 接收到查询请求 -  身份证: %s", req.IdCard)

	// 创建人员服务
	personService := service.NewPersonService()
	personSearch, err := personService.GetPerson(req.IdCard, "", "")
	if err != nil {
		utils.BadRequest(c, "查询失败: "+err.Error())
		log.Printf("[SearchPerson] 查询失败: %v", err)
		return
	}

	if personSearch.PersonInfo == nil {
		utils.NotFound(c, "未找到人员信息")
		log.Printf("[SearchPerson] 未找到人员信息 - 身份证: %s", req.IdCard)
		return
	}

	//log.Printf("[SearchPerson] 查询成功 - 身份证: %s, 姓名: %s", personSearch.PersonInfo.IdCard, personSearch.PersonInfo.Name)

	// 整合所有返回数据到统一结构
	responseData := struct {
		*model.PersonSearch
		Query string `json:"query"`
		Type  string `json:"type"`
	}{PersonSearch: personSearch, Query: req.IdCard, Type: "json"}

	utils.Success(c, gin.H{
		"person": responseData,
	})
}

// GetPersonProfile 获取人员画像
func GetPersonProfile(c *gin.Context) {
	id := c.Param("id")
	utils.Success(c, gin.H{
		"person_id":        id,
		"basic_info":       gin.H{},
		"police_records":   []gin.H{},
		"social_relations": []gin.H{},
		"social_data":      []gin.H{},
		"internal_data":    []gin.H{},
	})
}

// GetPersonRelations 获取人员社会关系
func GetPersonRelations(c *gin.Context) {
	id := c.Param("id")
	utils.Success(c, gin.H{
		"person_id": id,
		"relations": []gin.H{},
	})
}

// GetPersonCases 获取人员案件记录
func GetPersonCases(c *gin.Context) {
	id := c.Param("id")
	utils.Success(c, gin.H{
		"person_id": id,
		"cases":     []gin.H{},
	})
}

// GetPersonList 获取人员基本信息集合
func GetPersonList(c *gin.Context) {
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
	personService := service.NewPersonService()
	personList, total, err := personService.GetPersonList(paramsStr, pageInt, pageSizeInt)
	if err != nil {
		utils.BadRequest(c, "获取人员列表失败: "+err.Error())
		return
	}

	// 返回结果
	utils.Success(c, gin.H{
		"person_list": personList,
		"total":       total,
		"page":        pageInt,
		"pageSize":    pageSizeInt,
	})
}
