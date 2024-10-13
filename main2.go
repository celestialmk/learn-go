package main

import (
	"fmt"
)

func updateName(x *string) { //pass pointers to function
	*x = "wedge" //dereference using asterisk
}

func main() {
	name := "martin"

	fmt.Println("memory address (memory block) of name is ", &name)

	// *string, means pointer to a string

	m := &name
	fmt.Println("memory address of m is ", m)
	fmt.Println("value at memory address ", *m)

	updateName(m)

	fmt.Println("Name is ", name)

}
