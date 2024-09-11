// Deep-Go is a sample library that demonstrates how to create a library in Go
// This is my first Go library
// Thanks for visiting my GitHub repository
package deep_go

import (
	"fmt"
	"log"
	"strings"
)

// Sum is a sample function that adds two integers
// For example, Sum(2, 3) returns 5
func Sum(a, b int) int {
	log.Println("You're using Deep-Go library")
	fmt.Println(strings.Repeat("=",50))
	return a + b
}

// Multiply is a sample function that multiplies two integers
// For example, Multiply(2, 3) returns 6
func Multiply(a, b int) int {
	log.Println("You're using Deep-Go library")
	fmt.Println(strings.Repeat("=",50))
	return a * b
}
