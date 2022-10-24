package regexp

import (
	"fmt"
	"testing"
)

func ExampleIsGoodPwd() {
	pwd := "a1111111"
	fmt.Println(IsGoodPwd(pwd))

	pwd1 := "!a111111"
	fmt.Println(IsGoodPwd(pwd1))

	// output:
	// false 密码包含至少一位数字，字母和特殊字符,且长度8-16
	// true !a111111
}

func TestIsEmail(t *testing.T) {

	_is := IsEmail("ministor@126.com")

	if !_is {
		t.Errorf("Isn't Mail")
	}

	_is = IsEmail("ministor11@126.cn.com")

	if !_is {
		t.Errorf("Isn't Mail")
	}

	_is = IsEmail("minis@tor11@126.cn.com")

	if _is {
		t.Errorf("Isn't Mail")
	}
}

func TestIsMobile(t *testing.T) {

	_is := IsMobile("13810167616")

	if !_is {
		t.Errorf("Isn't Mobile")
	}

}

func TestMobileReplaceRepl(t *testing.T) {

	_str := MobileReplaceRepl("13910067616")

	fmt.Println(_str)
	if _str != "139****7616" {
		t.Errorf("MobileReplaceRepl")
	}
}
