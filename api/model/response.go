package model

const (
	PARAMATER_VALIDATION_ERROR  = "参数不合法"
	UNKNOWN_ERROR="未知错误"

)

// ReponseCommonModel 响应通用模型
type ReponseCommonModel struct{
	Status string
	Data interface{}
}

// ReponseErrModel 当发生错误时，响应的错误模型
type ReponseErrModel struct{
	ErrCode string
	ErrMessage string
}

