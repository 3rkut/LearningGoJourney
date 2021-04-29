package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int { // paramater is important in this example.
	fmt.Println("p : ", p) // for better understanding.
	fmt.Println("*p : ", *p)
	fmt.Println("&p : ", &p)
	*p++ // v += 1 at this point.
	return *p
}

func main() {

	fmt.Println(f() == f()) // Prints false.

	v := 1                            // v = 1.
	fmt.Println("address of v: ", &v) // check v's address.
	incr(&v)                          // parameter > address of v.
	// now v is 2.
	fmt.Println("new v value: ", v)   // check v's value.
	fmt.Println(incr(&v))             // now v is 3.
	fmt.Println("address of v: ", &v) // still same address.
}
