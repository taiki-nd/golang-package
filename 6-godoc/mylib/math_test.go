package mylib

import (
	"fmt"
	"testing"
)

//Example

func Example() {
	vv := Average([]int{1, 2, 3, 4, 5})
	fmt.Println(vv)
}
func ExampleAverage() {
	v := Average([]int{1, 2, 3, 4, 5})
	fmt.Println(v)
}

func ExamplePerson2_Say() {
	p := person2{"mike", 20}
	p.Say()
}

func TestAverage(t *testing.T) {
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}
