package main

import "fmt"

const aConst int = 32

const implicitConst = 67

func main(){

	// Practice variables, types and constants printing

	var aString string ="this is Go"
	fmt.Println(aString)
	fmt.Printf("the variable type is %T\n", aString)

	var anInteger int =43
	fmt.Println(anInteger)

	var defaultInt int
	// Prints 0
	fmt.Println(defaultInt)

	var defaultString string
	// Prints empty string
	fmt.Println(defaultString)

	var implicitInt = 52
	// Prints 52
	fmt.Println(implicitInt)

	var implicitString = "This is implicityl a string"
	fmt.Println(implicitString)

	// This will only work inside a function
	myString := "This is string without var"
	fmt.Println(myString)

	fmt.Println(aConst)
	fmt.Println(implicitConst)

	
}