// lecture0
package main

import (
	"fmt"
)

/*
func deneme(x, y int) (int, int, int) {
    return x + y, x * y, x - y
}

func yazdir(isim string, soyisim string, yas int) (string, string,int ) {
    return "Hello " + isim +"\n", "Your surname is:" + soyisim + "\n", yas
}*/

func main() {

	type Person struct {
		Name string
		Age  int
	}

	insan1 := Person{
		Name: "Adam",
		Age:  32,
	}

	insan2 := Person{Name: "Ahmet", Age: 22}
	insan2.Age = 32

	insan3 := Person{"SADASDA", 40}

	fmt.Printf("insan1 ad: %d \ninsan1 yas: %d \ninsan2 ad: %d \ninsan2 yas: %d \ninsan3 ad: %d \ninsan3 yas: %d", insan1.Name, insan1.Age, insan2.Name, insan2.Age, insan3.Name, insan3.Age)

	fmt.Println("Isim: " + insan1.Name)
	fmt.Printf("insan1'in yasi: %d\n", insan1.Age)
	var ( // daha clear bi syntax.
		xddd1 = "selam"
		xddd2 = 3.14
	)
	fmt.Printf("xddd1 tipi: %T\nxddd2 tipi: %T\n", xddd1, xddd2)

	/*
	   var myvar1,myvar2,myvar3 = deneme(12,7)
	   fmt.Printf("Toplama: %d\n", myvar1)
	   fmt.Printf("Carpma: %d\n", myvar2)
	   fmt.Printf("Cikarma: %d\n", myvar3)
	   deneme := "bitti"
	   fmt.Printf("%s\n", deneme)
	   fmt.Println(yazdir("adam", "snow", 43))

	   a1 := 3000
	   fmt.Printf("a1: %d", a1)*/

}
