package main

import (
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)        //change to uppercase
	names := strings.Split(s, " ") //create an array while splitting at open space

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}

func main() {
	// name := "Martin"
	// age := 25
	// fmt.Printf("my name is %v and my age is %v \n", name, age)
	// var SprintStr = fmt.Sprintf("my name is %v and my age is %v \n", name, age)
	// fmt.Println("Saved string", SprintStr)

	// //%T
	// //%0.1f

	// //Arrays [length of array]type{contents}
	// var ages = [3]int{20, 40, 50}
	// names := [4]string{"Martin", "Michael", "Marshal", "Malcom"}

	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// fruits := []string{"apple", "banana", "cherry", "mango"}

	// //loops
	// x := 0

	// //while loop
	// for x < 6 {
	// 	fmt.Println("Value of x is: ", x)
	// 	x++
	// }

	// //modified while
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of i is:", i)
	// }

	// //better for loop
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Println(fruits[i], "is a fruit")
	// }

	// //Using for in
	// for index, value := range fruits {
	// 	fmt.Printf("the position of %v is %v \n", value, index)
	// }

	// //Using for in without index
	// for _, value := range fruits {
	// 	fmt.Printf("the value is %v \n", value)
	// }

	// new_age := 30

	// if new_age < 40 {
	// 	fmt.Println("age is less than 40")
	// }

	// fn1, fn2 := getInitials("Martin Ssemakula")
	// println(fn1, fn2)

	// fn3, fn4 := getInitials("Marshal")
	// println(fn3, fn4)

	//maps
	// menu := map[string]float64{ //specify type of key and value
	// 	"beans": 3.55,
	// 	"meat":  4.99,
	// }

	// fmt.Println(menu["beans"])

	// for k, v := range menu {
	// 	fmt.Println(k, "-", v)
	// }

	//In a single map, all keys are the same type, all values are the same type

	//get value at memory address .... *m
	//get memory address ...&m

	//Pointers allow to update value of datatype in  functions..kinda directly

}
