// Lesson 1
// This is a simple Go program that prints "Hello From Go!" to the console.

package main // The package declaration. In Go, every file belongs to a package. The "main" package is a special package that indicates that this file is an executable program.

import "fmt" // Importing the "fmt" package to use the Println function for output.

func main() { // The main function is the entry point of the program. When the program is executed, this function is called.

	fmt.Println("Hello From Go!")
}

//Output:
//Command: go run main.go
//Hello From Go!

//Binary to Go Code:
//Command: go build main.go
//./main
//Hello From Go!
