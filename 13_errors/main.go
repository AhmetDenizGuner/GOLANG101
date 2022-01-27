package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {

	//GO hatalara uygulamanın bize verdiği bir değer olarak bakar.
	// nil başlangıç değeri olmayan ifadelere verilen değerdir.

	result, err := evenNum(11)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Girdiğiniz sayı: ", result)
	}

	//***************************************

	//error interface türünden bir typedir.

	result1, err1 := sRoot(-5)
	if err != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}

	//**********************************************

	file, err9 := os.Open("test.txt")
	if err9 != nil {
		fmt.Println(err9)
	} else {
		fmt.Println("Dosyamız", file)
	}

}

func evenNum(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("HATA: Çift sayı girmediniz")
	}
	return num, nil
}

func sRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Negatif sayıların karekökü alınamaz")
	}
	return math.Sqrt(num), nil
}

/* package main
import (
	"fmt"
	"math"
)
func main() {
	result := sRoot(-4)
	{
		fmt.Println(result)
	}
}
func sRoot(num float64) float64 {
	return math.Sqrt(num)
} */ //NaN döner,beklenmeyen sonuç bir hatadır.
