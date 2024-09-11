package deep_go

import (
	"fmt"
	"log"
	"strings"
)

func Sum(a, b int) int {
	log.Println("You're using Deep-Go library")
	fmt.Println(strings.Repeat("=",50))
	return a + b
}

func Multiply(a, b int) int {
	log.Println("You're using Deep-Go library")
	fmt.Println(strings.Repeat("=",50))
	return a * b
}
