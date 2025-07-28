package model

// PersonFile 人员档案模型
// 可根据实际需求扩展字段
// 这里只包含基本信息

type PersonFile struct {
	// 查询类型1：通过手机号或姓名查基本信息
	PersonInfo *Person `json:"person_info,omitempty"`
	// 查询类型2：通过身份证号查手机号和部门
	Phone       string     `json:"phone,omitempty"`
	Department  string     `json:"department,omitempty"`
	Tags        []string   `json:"tags,omitempty"`
	MedicalInfo []*Medical `json:"medical_info,omitempty"`
}
