// types/constants/begin/main.go
package main

import (
	"fmt"
	"unicode"
)

// declare a constant representing pi
const pi = 3.1419
const a rune = 'a'

// declare a rune constant for the letter 'a'

func main() {
	// print the value and types of pi and a
	fmt.Printf("pi: %v - %T\n", pi, pi)
	fmt.Printf("a: %c - %T\n", a, a)

	// use unicode package to confirm that the rune is a letter
	fmt.Print(unicode.IsLetter(a))
}
