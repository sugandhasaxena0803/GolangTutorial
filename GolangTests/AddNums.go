package GolangTests

func AddNums(sl ...int) int {
	sum := 0
	for _, v := range sl {
		sum = sum + v
	}
	return sum
}
