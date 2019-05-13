package kit

import "regexp"

// Validate 全局的验证变量
var Validate = validate{}

type validate struct{}

// IsPhone 验证是否是电话
func (v validate) IsPhone(phone string) bool {
	r := regexp.MustCompile(`^1\d{10}$`)
	return r.MatchString(phone)
}
