package main

import "fmt"

func main() {

	var name string = "Ahmet"
	//or
	var name2 string
	name2 = "Deniz"

	var age int = 24

	var isMarried bool
	isMarried = false

	//var firstName, secondName string = "Ahmet", "Deniz"

	age2 := 25 // short way değişken tanımlama

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(age)
	fmt.Println(isMarried)

	fmt.Printf("%T\n", age2)

}
