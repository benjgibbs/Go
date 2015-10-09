package main

import (
	"fmt"
)

type StringPrinter interface {
	printString(s string)
}

type IntPrinter interface {
	printInt(s int)
}

type Printer struct {
	prefix string
}

func (p Printer) printInt(s int) {
	fmt.Printf("%s: %d\n", p.prefix, s)
}

func (p Printer) printString(s string) {
	fmt.Printf("%s: %s\n", p.prefix, s)
}

func callStringPrinter(sp StringPrinter, s string) {
	sp.printString(s)
}

func callIntPrinter(ip IntPrinter, i int) {
	ip.printInt(i)
}

//func (s string) enquoteExtensionMethod() string {
//	return "'" + s + "'"
//}

func main() {
	fmt.Println("Hello")
	p := Printer{"Printer"}
	callIntPrinter(p, 101)
	callStringPrinter(p, "Test")
}
