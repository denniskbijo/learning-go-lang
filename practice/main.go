package main

import "fmt"

func main(){
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

	var implicitInt int = 52
	// Prints 52
	fmt.Println(implicitInt)

	var implicitString = "This is implicityl a string"
	fmt.Println(implicitString)
	
}