//Lesson 3: Integers (Signed, Unsigned, Bytes, and Runes)

package main

import "fmt"

func main() {
	// Variable of type uint8
	var smallPositiveValue uint8
	smallPositiveValue = 10
	fmt.Println(smallPositiveValue)

	// Variable of type int8
	var smallPositiveNegativeInt int8
	smallPositiveNegativeInt = -10
	fmt.Println(smallPositiveNegativeInt)

	// uint16, uint32, uint64
	// int16, int32, int64
	var myInt int = 240000000
	fmt.Println(myInt)

	myInt = int(smallPositiveNegativeInt)
	myInt = int(smallPositiveValue)
	//to typecase a value int()

	var myByte byte = 'A'
	fmt.Println(myByte)
	//Byte is mainly use to represent raw data
	//and also to distinguish between uint8 and byte
	//since byte is an alias for uint8

	// Rune is an alias for int32
	var myRune rune = 'B'
	fmt.Println(myRune)
	myRune = '😁'
	fmt.Println(myRune)

}

//Output:
//Command: go run main.go
// 10
// -10
// 240000000
// 65
// 66
// 128513
