package BenchMark

import "strings"

func AddNumbers(n int) int {
	sum := 0
	for x := 0; x < n; x++ {
		sum = sum + x
	}

	return sum

}

func ConcatSlice(sl []string) string {
	s := ""
	for _, v := range sl {
		s = s + v
		s = s + " "
	}

	return s

}

func JoinSlice(sl []string) string {
	return strings.Join(sl, " ")
}
