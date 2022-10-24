package strings

import "fmt"

func ExampleStrByXOR() {
	var message = "hello world"

	var key = "2333"

	var encodeStr = StrByXOR(message, key)
	fmt.Println(encodeStr)

	var decodeStr = StrByXOR(encodeStr, key)
	fmt.Println(decodeStr)

	// output:
	// ZV__]D\@_W
	// hello world
}
