/*
mylib is special lib
*/
package mylib

//Average returns the average
func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}

//person2 doc
type person2 struct {
	Name string
	Age  int
}
