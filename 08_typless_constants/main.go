package main

import "fmt"

func main() {

	/* 	const x = 5
	   	fmt.Printf("%T, %v\n", x, x) */

	/* 	const x = 3
	   	var y int16 = 12
	   	fmt.Printf("%T, %v\n", x, x)
	   	fmt.Printf("%T, %v\n", y, y)
	   	fmt.Printf("%T, %v\n", x+y, x+y) // bu satır hata vermez otomatik convert etti ama gecici int16 oldu burada
		   fmt.Printf("%T, %v\n", x, x) */ //burda varsayılan int e döndü, ama eğer typed olsa hata veriridi elde convert etmek lazımdı

	/* 	const x = int16(5.2 + 4.8)
	   	fmt.Printf("%T, %v\n", x, x) */

	/* 	const x = 3
	   	const y = 5.6
	   	fmt.Printf("%T, %v\n", x, x)
	   	fmt.Printf("%T, %v\n", y, y)
	   	fmt.Printf("%T, %v\n", x+y, x+y) */ //Hata vermez typless constant x'i gecici floata convert eder

	/*var a = 3
	var b = 5.2
	fmt.Println(a + b) */ //ama bu hata verir, bu özellik sadece typless constantlara özel

	const x = 3
	const y = 5.6
	const z = true
	const t = "test"

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", t, t)

	//Sabitlerde de shadowing kavramı çalışır.

}
