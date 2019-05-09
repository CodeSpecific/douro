package model

// ResponseStatus 响应状态
type ResponseStatus string

const (
	SUCCESS ResponseStatus = "Success"
	FAIL="Fail"
)

// ResponseCommonModel 全局返回对象实例
var ResponseCommonModel = &responseCommonModel{}

// responseCommonModel 响应通用模型
type responseCommonModel struct {
	Status ResponseStatus
	Data   interface{}
}

// Success 当响应成功时，需要调取的方法
func (m *responseCommonModel) Success(data interface{}) *responseCommonModel {
	m.Status = SUCCESS
	m.Data = data
	return m
}

// Error 当响应失败时，需要调取的方法
func (m *responseCommonModel) Error(code, msg string) *responseCommonModel {
	m.Status = FAIL
	err := &reponseErrModel{code,msg}
	m.Data = err
	return m
}

// reponseErrModel 当发生错误时，响应的错误模型
type reponseErrModel struct {
	ErrCode string  `json:"Code"`
	ErrMsg  string	`json:"Message"`
}
