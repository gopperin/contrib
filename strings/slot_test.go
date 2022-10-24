package strings

import "fmt"

func ExampleSumASCII() {
	fmt.Println(SumASCII("ministor@126.com") % 10)

	fmt.Println(SumASCII("") % 10)

	fmt.Println(SumASCII("eric@126.com") % 10)

	fmt.Println(SumASCII("realestate@126.com") % 10)

	// output:
	// 7
	// 0
	// 1
	// 8
}
