package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good Morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func saysomething(h human) {
	h.speak()
}

func main() {
	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)

	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)

	p1 := person{
		"Miss",
		"Moneypenny",
	}
	fmt.Println(p1)
	// p1.speak()
	saysomething(p1)

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	// sa1.speak()
	saysomething(sa1)
}
