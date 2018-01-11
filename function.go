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

func Greet(salutation Salutation){
	fmt.Println(salutation.name)
	fmt.Println(salutation.greeting)
}

func main() {

	var s = Salutation{name: "Joe", greeting: "Hello!"}

	Greet(s)

}
