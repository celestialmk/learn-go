package main

import "fmt"


//constants cannot be declared with :=, := is only used inside functions
const Pi = 3.14

func main(){
	const name = "Marty"
	fmt.Println("Happy", Pi, "day", name)

}