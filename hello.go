package main

import "fmt"

type Salutation struct {
	name string
	greeting string
}

func main() {

	var s = Salutation{name: "Joe", greeting: "Hello!"}

	fmt.Println(s.name, s.greeting)
}
