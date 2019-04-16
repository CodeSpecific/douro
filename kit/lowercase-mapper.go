package kit

import "strings"

// LowerCaseMapper 实现Xorm的core.IMapper接口，实现自定义表名和结构体之间的映射
type LowerCaseMapper struct{}

// Obj2Table 实现 core.IMapper 接口的方法
func (m LowerCaseMapper) Obj2Table(o string) string {
	return strings.ToLower(o)
}

// Table2Obj 实现 core.IMapper 接口的方法
func (m LowerCaseMapper) Table2Obj(t string) string {
	return strings.ToUpper(t)
}
