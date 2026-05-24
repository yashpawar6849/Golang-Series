//Lesson 2: Data Types in Go

package main

import "fmt"

func main() {
	var mystring string // string is a data type that can hold a sequence of characters
	//String zero value is ""
	mystring = "Welcome To Golang Series"
	fmt.Println(mystring)

	var myInt int // int is a data type that can hold whole numbers
	//zero value for number types is 0
	myInt = 100
	fmt.Println(myInt)

	var myBool bool // bool is a data type that can hold true or false values
	//Zero value for boolean is false
	myBool = true
	fmt.Println(myBool)

}
