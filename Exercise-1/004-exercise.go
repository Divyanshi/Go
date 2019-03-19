package main

import "fmt"

type myType int

var myVar myType

func main() {
	fmt.Println(myVar)
	myVar = 45
	fmt.Println(myVar)
	fmt.Printf("%T\n", myVar)
}
