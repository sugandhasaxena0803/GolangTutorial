package BenchMark

import (
	"fmt"
	"strings"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	result := 0
	result = AddNumbers(5)
	if result != 10 {
		t.Error("Expected", 10, "got", result)
	}
}

func ExampleAddNumbers() {
	fmt.Println(AddNumbers(6))
	//Output:
	//15

}

func BenchmarkAddNumbers(b *testing.B) {

	for i := 0; i < b.N; i++ {
		AddNumbers(5)
	}

}

func BenchmarkConcatSlice(b *testing.B) {
	s := "Hello my name is sugandha. I am a very confident girl . I can do it "
	sl := strings.Split(s, " ")
	for i := 0; i < b.N; i++ {
		ConcatSlice(sl)
	}

}

func BenchmarkJoinSlice(b *testing.B) {
	s := "Hello my name is sugandha. I am a very confident girl . I can do it "
	sl := strings.Split(s, " ")
	for i := 0; i < b.N; i++ {
		JoinSlice(sl)
	}
}
