package main

import "fmt"

var packVar = "Package Scope"

//In package scope, you cant use short variable declaration.
//Because go compiler use keywords for starting compile.
//This variables hold place on memory while all run time

func main() {

	{
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
		var packVar = "Block Override Scope"
		fmt.Println(packVar)
		//bu scope da başka birşey atamam globaldeki değeri değiştirmez, dışarda yine package scope olarak print edildi.

	}

	var funcVar = "Func Scope"
	fmt.Println(funcVar)
	fmt.Println(packVar)

	packVar = "Block Override Scope"
	//ama var kullanmadan assigning yapınca globaldeki değeri değiştirmiş olduk ardık block override olarak set edildi.
	fmt.Println(packVar)
	{
		fmt.Println(packVar)
	}

	/*
		var name = "Arin"
		name := "Elis" hata verir ama
		name,surname := "Elis", "Soft" hata vermez!!!
	*/

}

//OUTPUT
/*
Block Scope
Block Override Scope
Func Scope
Package Scope
Block Override Scope
Block Override Scope
*/
