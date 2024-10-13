package main

import "fmt"

func main() {
	mybill := newBill("Martin's bill")

	fmt.Println(mybill)

	fmt.Println(mybill.format())
}
