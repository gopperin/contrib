package regexp

import (
	"regexp"

	"github.com/dlclark/regexp2"
)

func IsGoodPwd(str string) (bool, string) {
	expr := `^(?![0-9a-zA-Z]+$)(?![a-zA-Z!@#$%^&*]+$)(?![0-9!@#$%^&*]+$)[0-9A-Za-z!@#$%^&*]{8,16}$`
	reg, _ := regexp2.Compile(expr, 0)
	m, _ := reg.FindStringMatch(str)
	if m != nil {
		res := m.String()
		return true, res
	}
	return false, "密码包含至少一位数字，字母和特殊字符,且长度8-16"
}

// IsEmail 是否是email
func IsEmail(email string) bool {
	if email == "" {
		return false
	}
	ok, _ := regexp.MatchString(`^([a-zA-Z0-9]+[_|\_|\.|\-]?)*[_a-z\-A-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.|\-]?)*[a-zA-Z0-9\-]+\.[0-9a-zA-Z]{2,6}$`, email)
	return ok
}

// IsUsername 是否只包含数字, 字母 -, _
func IsUsername(username string) bool {
	if username == "" {
		return false
	}
	ok, _ := regexp.MatchString(`[^0-9a-zA-Z_\-]`, username)
	return !ok
}

const (
	regularMobile = "^1([38][0-9]|14[57]|5[^4])\\d{8}$"
)

// IsMobile check if mobile
func IsMobile(mobileNum string) bool {
	reg := regexp.MustCompile(regularMobile)
	return reg.MatchString(mobileNum)
}

// MobileReplaceRepl MobileReplaceRepl 手机号脱敏
func MobileReplaceRepl(str string) string {
	re, _ := regexp.Compile(`(\d{3})(\d{4})(\d{4})`)
	return re.ReplaceAllString(str, "$1****$3")
}
