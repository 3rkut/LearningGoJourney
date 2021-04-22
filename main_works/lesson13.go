package main

import "fmt"

type Ogrenci struct {
	yas    int
	isim   string
	notlar []int
}

func (x Ogrenci) getAge() int {
	return x.yas

}

func (y *Ogrenci) setAge(yas int) {
	y.yas = yas
}

func main() {

	ogrenci1 := Ogrenci{isim: "Case", yas: 20, notlar: []int{70, 80}}
	fmt.Println(ogrenci1.notlar)
	fmt.Println(ogrenci1.getAge())

	ogrenci2 := Ogrenci{isim: "Johny", yas: 18, notlar: []int{55, 60}}
	fmt.Println(ogrenci2)
	ogrenci2.setAge(19)
	fmt.Println(ogrenci2)

}
