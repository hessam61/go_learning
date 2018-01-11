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
	message, alternate := CreateMessage(salutation.name, salutation.greeting)
	fmt.Println(message)
	fmt.Println(alternate)
}

func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "HEY!"  + name
	return
}

func main() {

	var s = Salutation{name: "Joe", greeting: "Hello"}

	Greet(s)

}
