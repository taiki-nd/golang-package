package mylib

var Public string = "Public"

//var private string = "private"

type Person struct {
	Name string
	Age  int
}

func Human() string {
	return "hello world"
}
