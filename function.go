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
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "Howdy")
	fmt.Println(message)
	fmt.Println(alternate)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	fmt.Println(len(greeting))
	message = greeting[1] + " " + name
	alternate = "HEY!"  + name
	return
}

func main() {

	var s = Salutation{name: "Joe", greeting: "Hello"}

	Greet(s)

}
