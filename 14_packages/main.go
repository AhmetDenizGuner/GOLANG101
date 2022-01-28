package main

import (
	"bufio"
	"fmt"
	"getcelcius"
	"merhaba"
	"os"
	"strings"
)

func main() {

	fmt.Print("Lütfen Sınav Sonucunuzu Giriniz: ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') // _ blank identifier
	fmt.Println(value)

	//******************************************************

	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Count("animatrix", "a"))

	//*******************************************************

	fmt.Println(strings.ToUpper("Gopher"))

	merhaba.Hello()
	//kendi paketini src altına tanımla paket adında klasor içinede paket adıyla.go tanımla !!!!!!

	fmt.Print("Lütfen Celcius dereceyi giriniz: ")
	celcius, err := getcelcius.GetCelcius()
	if err != nil {
		fmt.Println(err)
	}

	kelvin := celcius + 273

	fmt.Println("Kelvin cinsinden sıcaklık:", kelvin)

}
