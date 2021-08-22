package auth

// VerifyStudentReqModel 学生认证：请求
type VerifyStudentReqModel struct {
}

// VerifyStudentResModel 学生认证：响应
type VerifyStudentResModel struct{}

// VerifyCompanyReqModel 公司认证：请求
type VerifyCompanyReqModel struct{}

// VerifyCompanyResModel 公司认证：响应
type VerifyCompanyResModel struct{}

// VerifyUniversityReqModel 高校认证：请求
type VerifyUniversityReqModel struct{}

// VerifyUniversityResModel 高校认证：响应
type VerifyUniversityResModel struct{}

// VerifyInstitutionReqModel 机构认证：请求
type VerifyInstitutionReqModel struct{}

// VerifyInstitutionResModel 机构认证：响应
type VerifyInstitutionResModel struct{}

// UniversityStructReqModel 学校的系统结构
type UniversityStructReqModel struct {
	Departments []string `json:"departments" form:"departments"`

	Struct map[int][]SingleStruct `json:"struct" form:"struct"`

	Students []struct {
		Name   string `json:"name"`
		SID    string `json:"sid"`
		CardID string `json:"card_id"`

		Campus  string `json:"campus"`
		College string `json:"college"`
		Grade   uint8  `json:"grade"`
		Major   string `json:"major"`
		Class   string `json:"class"`
	} `json:"students" form:"students"`
}

type SingleStruct struct {
	Campus  string `json:"campus"`
	College string `json:"college"`
	Grade   uint8  `json:"grade"`
	Major   string `json:"major"`
	Class   string `json:"class"`

	UID uint `json:"uid"`
}
