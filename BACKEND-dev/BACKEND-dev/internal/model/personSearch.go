package model

// personSearch 人员模型
type PersonSearch struct {
	PersonInfo  []*Person  `json:"person_info"`
	MedicalInfo []*Medical `json:"medical_info"`
}
