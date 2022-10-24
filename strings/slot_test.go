package strings

import "fmt"

func ExampleSumASCII() {
	fmt.Println(SumASCII("ministor@126.com") % 10)

	fmt.Println(SumASCII("") % 10)

	// output:
	// 7
	// 0
}
