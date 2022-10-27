package auth

import (
	"fmt"
)

func ExampleGoogle() {

	secret := "B5T74E3MZERRJER2IWUB4Y7BU6DTYMWZ"

	// secret最好持久化保存在
	// 验证,动态码(从谷歌验证器获取或者freeotp获取)
	code := "943928"
	bool, err := New(6).VerifyCode(secret, code)
	if bool {
		fmt.Println("√")
	} else {
		fmt.Println("X", err)
	}

	// output:
	// X <nil>
}

func ExampleDefault() {

	// 秘钥
	// secret := Default().GetSecret()
	// fmt.Println("Secret:", secret)

	// output:
}
