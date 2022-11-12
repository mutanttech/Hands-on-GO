// types/structs/fields/begin/main.go
package main

import "fmt"

// define a struct type for author
type author struct {
		firstname, lastname string
	}
//

func main() {
	// intialize author
	a := author {
		firstname: "Sam",
		lastname: "Diesel",
	} 
	//

	// print the author
	 fmt.Printf("%#v\n", a)
}
