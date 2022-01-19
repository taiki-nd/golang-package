package main

import (
	"fmt"
	"package/1-package/mylib"
	"package/1-package/mylib/sublib"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	fmt.Println(mylib.Human())

	sublib.Say()
}
