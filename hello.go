package main

import "fmt"

type Salutation struct {
	name string
	greeting string
}

const(
	PI = 3.14
	Language = "Go"
)
const (
	A = iota
	B = iota
	C = iota
)

func main() {

	var s = Salutation{name: "Joe", greeting: "Hello!"}

	fmt.Println(s.name, s.greeting)
	fmt.Println(PI, Language)
	fmt.Println(A, B, C)
}
