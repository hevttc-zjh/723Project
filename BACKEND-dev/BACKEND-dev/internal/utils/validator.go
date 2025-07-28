package utils

import (
	"regexp"
	"strings"
)

// ValidateIDCard 校验大陆居民身份证号
// 支持18位身份证号，包含最后一位可能的X/x
func ValidateIDCard(idCard string) bool {
	idCard = strings.ToUpper(strings.TrimSpace(idCard))
	if len(idCard) != 18 {
		return false
	}

	// 前17位必须是数字
	re := regexp.MustCompile(`^\d{17}[\dX]$`)
	if !re.MatchString(idCard) {
		return false
	}

	// 校验码计算（可选增强）
	// 此处可添加更严格的校验码算法，示例暂时省略
	return true
}

// ExtractIDCard 从字符串中提取18位身份证号
func ExtractIDCard(s string) string {
	s = strings.TrimSpace(s)
	re := regexp.MustCompile(`\b\d{17}[\dXx]\b`)
	matches := re.FindStringSubmatch(s)
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}

// ExtractMobile 从字符串中提取11位手机号
func ExtractMobile(s string) string {
	s = strings.TrimSpace(s)
	re := regexp.MustCompile(`\b1(3[0-9]|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8[0-9]|9[0-35-9])\d{8}\b`)
	matches := re.FindStringSubmatch(s)
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}

// ExtractName 从字符串中提取姓名 允许2-6个汉字
func ExtractName(s string) string {
	s = strings.TrimSpace(s)
	re := regexp.MustCompile(`^[\p{Han}]{2,6}$`)
	matches := re.FindStringSubmatch(s)
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}

// ValidateMobile 校验三大运营商手机号
// 支持移动、联通、电信最新号段（截至2024年）
func ValidateMobile(mobile string) bool {
	mobile = strings.TrimSpace(mobile)
	if len(mobile) != 11 {
		return false
	}

	// 匹配三大运营商号段正则
	pattern := `^1(3[0-9]|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8[0-9]|9[0-35-9])\d{8}$`
	match, _ := regexp.MatchString(pattern, mobile)
	if !match {
		return false
	}

	// 进一步区分运营商（可选）
	// mobileOperator := getMobileOperator(mobile)
	return true
}

// 可选：获取运营商类型（示例）
// func getMobileOperator(mobile string) string {
// 	prefix := mobile[:3]
//  	switch {
//  	case contains([]string{"134","135","136","137","138","139","150","151","152","157","158","159","182","183","184","187","188"}, prefix):
//  		return "移动"
//  	case contains([]string{"130","131","132","155","156","185","186","166","175","176"}, prefix):
//  		return "联通"
//  	case contains([]string{"133","153","173","177","180","181","189","199"}, prefix):
//  		return "电信"
//  	default:
//  		return "未知"
//  	}
// }`

// contains 辅助函数（示例）
// func contains(slice []string, item string) bool {
//  	for _, s := range slice {
//  		if s == item {
//  			return true
//  		}
//  	}
//  	return false
// }`

// 校验名字
func ValidateName(s string) bool {
	// 允许2-6个汉字
	match, _ := regexp.MatchString(`^[\u4e00-\u9fa5]{2,6}$`, s)
	return match
}
