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

func Greet(salutation Salutation, do func(string)){
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "Howdy")
	do(message)
	do(alternate)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	fmt.Println(len(greeting))
	message = greeting[1] + " " + name
	alternate = "HEY!"  + name
	return
}

func Print(s string){
	fmt.Print(s)
}

func Println(s string){
	fmt.Println(s)
}

func main() {

	var s = Salutation{name: "Joe", greeting: "Hello"}

	Greet(s, Println)

}
