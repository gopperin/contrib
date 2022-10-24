package strings

import "strings"

func SumASCII(src string) int64 {
	tmp := strings.ToLower(src)
	var ret int64
	for i := 0; i < len(tmp); i++ {
		ret += int64(rune(tmp[i]))
	}
	return ret
}
