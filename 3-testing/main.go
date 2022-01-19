package main

import (
	"fmt"
	"package/3-testing/mylib"
	"package/3-testing/mylib/sublib"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	fmt.Println(mylib.Human())
	sublib.Say()

	person := mylib.Person{Name: "Mike", Age: 20}
	fmt.Println(person)

	fmt.Println(mylib.Public)
	//fmt.Println(mylib.private)

}
