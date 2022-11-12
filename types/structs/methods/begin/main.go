// types/structs/methods/begin/main.go
package main

import "fmt"

type author struct {
	first  string
	last   string
	gender bool
}

// fullName returns the full name of the author
func fullName(a author) string {
	message := ""
	if a.gender {
		message = a.first + " " + a.last + " is Male."
	} else {
		message = a.first + " " + a.last + " is Female."
	}
	return message
}

//

func main() {
	// initialize author
	a := author{
		first:  "Angelina",
		last:   "Jolie",
		gender: false,
	}

	// print the author's full name
	fmt.Println(fullName(a))
	//

	b := author{
		first:  "Brad",
		last:   "Pitt",
		gender: true,
	}

	// print the author's full name
	fmt.Println(fullName(b))
	//
}
