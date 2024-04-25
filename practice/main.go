package main

import (
	// "bufio"
	"fmt"
	"math"
	// "os"
	// "strconv"
	// "strings"
)

// const aConst int = 32

// const implicitConst = 67

func main() {

	// Practice variables, types and constants printing

	// var aString string = "this is Go"
	// fmt.Println(aString)
	// fmt.Printf("the variable type is %T\n", aString)

	// var anInteger int = 43
	// fmt.Println(anInteger)

	// var defaultInt int
	// // Prints 0
	// fmt.Println(defaultInt)

	// var defaultString string
	// // Prints empty string
	// fmt.Println(defaultString)

	// var implicitInt = 52
	// // Prints 52
	// fmt.Println(implicitInt)

	// var implicitString = "This is implicityl a string"
	// fmt.Println(implicitString)

	// // This will only work inside a function
	// myString := "This is string without var"
	// fmt.Println(myString)

	// fmt.Println(aConst)
	// fmt.Println(implicitConst)

	// // Practice getting input from console

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter Text: ")

	// //Var _ is a variable we can ignore as it is used to handle error here
	// input, _ := reader.ReadString('\n')
	// fmt.Println("You entered: ", input)

	// //Convert input to float type
	// fmt.Print("Enter a number: ")

	// //Var err is an error object used to handle error in conversion here
	// numInput, _ := reader.ReadString('\n')
	// fmt.Println("You entered: ", numInput)

	// afloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Parsed number is: ", afloat)
	// }

	//Math Package

	//Multiple var declaration
	i1, i2, i3 := 12, 13, 15
	intSum := i1 + i2 + i3
	fmt.Println(intSum)

	f1, f2, f3 := 23.5, 64.1, 76.3
	floatSum := f1 + f2 + f3
	//Prints 163.89999999999998
	fmt.Println(floatSum)

	//Safely Rounding to keep float precision
	//Multiply by 100 before rounding and divide by 100 after rounding
	floatSum = math.Round(floatSum*100) / 100
	//Prints 163.9
	fmt.Println("The sum is now: ", floatSum)
}
