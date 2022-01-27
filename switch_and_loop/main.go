package main

import "fmt"

func main() {

	// Switch

	grade := -7
	switch grade { //ifde olduğu gibi --> switch grade := 3 {} yapılabilir
	default:
		fmt.Println("Geçersiz Not")
	case 5: // if grade == 5 { fmt.Println("Pekiyi")  }
		fmt.Println("Pekiyi")
	case 4:
		fmt.Println("İyi")
	case 3:
		fmt.Println("Orta")
	case 2:
		fmt.Println("Geçer")
		y := 100
		fmt.Println(y)
	case 1:
		fmt.Println("Başarısız")
	}

	//hicbir case e girmezse defaulte girer
	//yine veri tipi önemlidir.
	//break manuel olarak yazılmaz varmış gibi çalışır.

	switch {
	case false:
		fmt.Println("Bu yazdığımız konsolda görünmez.")

	case true:
		fmt.Println("Bu yazdığımız konsolda görünecek.")
	case true:
		fmt.Println("buda görünmez") //cünku hayali break var

	}

	//LOOPS
	//sadece for loop var while yok ama for kullanılarak while yaratılabilir
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	//For is Go's "while"
	fmt.Println(sum)

	/* 	for i := 0; i <= 10; i += 5 {
		fmt.Println(i)
	} */

	// for <init statement>; <condition>; <post statement> { ----- }

	/* 	i := 0
	   	for ; i <= 10; i += 5 {
	   		fmt.Println(i)
	   	}
	   	fmt.Println(i) */

	/* 	for {  // Infinite Loop
		fmt.Println("Benim Adım Arin")
	} */

	/* 	for i := 0; true; i += 5 {
		fmt.Println(i)
	} */

	/* 	for i := 0; false; i += 5 {
		fmt.Println(i)
	} */

	/* 	i := 10
	   	for i >= 0 {
	   		fmt.Println(i)
	   		i--
	   	} */

	/* 	for i := 0; i <= 10; i++ { // continue --> döngünün başına git
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	} */

	for i := 0; i <= 10; i++ {

		if i == 3 {
			break
		}

		fmt.Println(i)
	} // break --> döngüden çık
	//continue o stepi geç

	//UYGULAMA

	//falltrough keywordu --> break keywordünü pas geçer

	/*switch x := 75; {
	case x < 20:
		fmt.Printf("%d küçüktür 20\n", x)
		fallthrough
	case x < 50:
		fmt.Printf("%d küçüktür 50\n", x)
		fallthrough
	case x < 100:
		fmt.Printf("%d küçüktür 100\n", x)
		fallthrough
	case x < 200:
		fmt.Printf("%d küçüktür 200\n", x)
	}*/

	//hem 100 hem de 200 bastırılır.

	//idiomatic kod yazmak GO'da parçaların bağımsız sade okunabilir yazılmasıdır.

	//less idiomatic
	if x := 20; x%2 == 0 {
		fmt.Println(x, "çifttir")
	} else {
		fmt.Println(x, "tektir")
	}

	//more idiomatic
	x := 20
	if x%2 == 0 {
		fmt.Println(x, "çifttir")
		return
	}
	fmt.Println(x, "tektir")

}

// 5-) 1 ile 50 arasındaki asal sayıları gösteren bir program yazınız.

func asal() {

	var x, y int

	for x = 2; x < 50; x++ {
		for y = 2; y < (x / y); y++ {
			if x%y == 0 {
				break
			}
		}

		if y > (x / y) {
			fmt.Printf("%d bir asal sayıdır\n", x)
		}
	}
}
