package main

import "fmt"

type myType int

var myVar myType
var anotherVar int

func main() {
	fmt.Println(myVar)
	myVar = 45
	fmt.Println(myVar)
	fmt.Printf("%T\n", myVar)
	anotherVar = int(myVar)
	fmt.Println(anotherVar)
	fmt.Printf("%T\n", anotherVar)

}
