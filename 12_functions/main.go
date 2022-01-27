package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	// Moduler Programming
	// Readable code
	// From Complex To Simple

	// func funcName(params) return type { code  }

	merhaba()

	// return vs print

	/* 	z := sum(5, 10)
	   	fmt.Println(z)
		   sum2(6, 11) */

	merhaba2("Elis", 4)

	// Fonksiyon İsimlendirme
	// ilk karekter harf
	// camel Case -- mySum, myBestFunction
	// paket dışında da kullancaksak o zaman ilk harf büyüktür aynısı değişiklikler içinde geçerli --> ilk harf büyük

	bölüm, kalan := bölme(104, 5)
	fmt.Println(bölüm, kalan)
	//Multiple return mümkün

	//merhaba("Arin", 6) Argument Function Run
	//func merhaba(name string, age int)  Parametre Function Write

	//farklı paketlere ve veri tiplerine ait olan fonksiyonlara metod denir.
	var x int = 10
	fmt.Println(x)
	var moment time.Time = time.Now() // Now ---> method
	fmt.Println(moment)

	fmt.Print("Lütfen Sınav Sonucunuzu Giriniz: ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') // _ blank identifier virgül olması sebebi ikinci değer hatayı mapledeği değişken
	// _ ile atanan değeri kullanmamamıza rağmen hata gelmedi,"\n" kullanamayız kuraldır
	fmt.Println(value)

	a := 5
	change(a)
	fmt.Println(a) //5 bastı C'deki gibi call by value

}

func change(a int) {
	a = 10
}

func sum(x, y int) int { //iki parametre aynı tipse sonuncuya yaz
	return x + y
}

func sum2(x, y int) {
	fmt.Println(x + y)
}

func merhaba() {
	fmt.Println("Benim Adım Arin")
}

func merhaba2(name string, age int) {
	fmt.Printf("Adım %s, yaşım %d", name, age)
}

func bölme(bölünen, bölen int) (bölüm, kalan int) {
	bölüm = bölünen / bölen
	kalan = bölünen % bölen

	return bölüm, kalan
}

/* package main
import "fmt"
func main() {
	fmt.Println(result(40))
}
func result(grade int) string {
	if grade >= 50 {
		return "geçtiniz"
	}
	return "kaldınız"
	fmt.Println("FONKSIYON ICINDEYIM") //bu satır hata verir erişelemez kod
} */
