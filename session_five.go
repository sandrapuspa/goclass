package main

import (
	// "reflect"
	"fmt"
)

type Employee struct {
	Human
	Subject string
}

type Human struct {
	Name string
	Address string
}

type Cast struct {
	Name string
	Debut string
}

type InterfaceMethod interface {
	SayHi()
	SayHellow()
	SayHay()
}

func SessionFive() {
	// h := Human{}
	// h.Name = "nichole"
	// h.Address = "NY"

	// c := ClassMate{}
	// c.Human = h
	// c.Subject = "biology"

	// // SayHi("nichole", "NY")
	// SayHi(1, 2)
	// h.SayHi()
	// c.SayHi()
	// var number float64 = 3.4

// 	upin := Employee{Human{"Upin", "123-456-789"}, "Google Inc"}
// 	var intf InterfaceMethod 
// 	intf = upin

// point := reflect.ValueOf(intf)
// // value := point.Elem()
// // value.SetFloat(7.1)
// intf.SayHellow()
// intf.SayHi()

// fmt.Println(point, "===")//, value)


}

// func SayHi(name, address string) {
// 	fmt.Printf("Hi %s, is %s you address?\n", name, address)
// }

func (h Human) SayHi(){//odd, even int) {
	// fmt.Printf("%d is odd and %d is even\n", odd, even)
	fmt.Println("Hellow")
}

func (h Human) SayHellow(){//odd, even int) {
	// fmt.Printf("Helooowwww~~ %d is odd and %d is even\n", odd, even)
	fmt.Println("World")
}

// func (c ClassMate) SayHi() {
// 	fmt.Printf("Hi there, my name is %s\n", c.Human.Name)
// }

func (c Cast) SayHay() {
	fmt.Printf("Hi there, my name is %s-%s\n", c.Name, c.Name)
}
