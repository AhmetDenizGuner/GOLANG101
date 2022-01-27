package main

import "fmt"

func main() {

	x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	/*
		10 int
		10 float64
	*/

	/*
		fmt.Println(x + y)
		\main.go:16:16: invalid operation: x + y (mismatched types int and float64)
	*/
	//Even both value are numbers, they are diffrent type so you cant operate them. int8 ve int16 bile toplanamaz.

	fmt.Println(x + int(y))
	fmt.Println(y + float64(x))
	//x or y'nin tipi değişmedi geçici bir değer oluştu
	//data loss'u önlemek için geniş değere dönüştürmek tercih edilmelidir.

	var x1 int16 = 128
	var y1 int8
	y1 = int8(x1)

	fmt.Println(y1) // -128 bastı

	//Stringler numbersa convertion yöntemiyle dönüştürülmez
	a := "100"
	b := 106
	fmt.Println(a + string(b))     //100☺ --> string içine int verince asci tablosundaki değeri getiriyor saçma bir değer geliyor.
	fmt.Println(a + fmt.Sprint(b)) //1001

}
