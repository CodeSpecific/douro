package kit

import "regexp"

// Validate 全局的验证变量,使用初始化形式，防止被外部修改
var Validate validate

func init() {
	Validate = validate{}
}

type validate struct{}

// IsPhone 验证是否是电话
func (v validate) IsPhone(phone string) bool {
	r := regexp.MustCompile(`^1\d{10}$`)
	return r.MatchString(phone)
}
