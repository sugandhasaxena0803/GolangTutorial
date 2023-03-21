// Package GolangTests  This is documentation ..adds nums
package GolangTests

// AddNums This is documentation ..adds nums
func AddNums(sl ...int) int {
	sum := 0
	for _, v := range sl {
		sum = sum + v
	}
	return sum
}
