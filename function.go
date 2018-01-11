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
	fmt.Println(CreateMessage(salutation.name, salutation.greeting))
}

func CreateMessage(name, greeting string) string {
	return greeting + " " + name
}

func main() {

	var s = Salutation{name: "Joe", greeting: "Hello"}

	Greet(s)

}
