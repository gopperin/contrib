package strings

func SumASCII(src string) int64 {

	var ret int64
	for i := 0; i < len(src); i++ {
		ret += int64(rune(src[i]))
	}
	return ret
}
