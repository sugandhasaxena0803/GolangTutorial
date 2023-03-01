package GolangTests

import "testing"

/*

ends with _test.go
in same package
func Testxxx(t *Testing

go test
go test -v
)
*/

func TestAddNums(t *testing.T) {
	sl := []int{9, 5, 1}
	n1 := AddNums(6, 8, 2)
	if n1 != 16 {
		t.Error("Expected ", 16, "Got", n1)
	}
	n1 = AddNums(sl...)
	if n1 != 15 {
		t.Error("Expected ", 15, "Got", n1)
	}

	//table tests
	type StructAddNum struct {
		input  []int
		output int
	}
	// thi is slice of type struct
	StructsAddNum := []StructAddNum{StructAddNum{[]int{9, 5, 1}, 15}, StructAddNum{[]int{6, 8, 2}, 16},
		StructAddNum{[]int{-9, 5, 4}, 0}}
	for _, vv := range StructsAddNum {
		n1 = AddNums(vv.input...)
		if n1 != vv.output {
			t.Error("Expected", vv.output, "got", n1)
		}

	}
}
