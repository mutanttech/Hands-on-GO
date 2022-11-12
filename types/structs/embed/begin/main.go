// types/structs/embed/begin/main.go
package main

import "fmt"

type person struct {
	first string
	last  string
}

// fullName returns the full name of a person
func (p person) fullName() string {
	return p.first + " " + p.last
}

// fullName returns the full name of a person
func (a author) fullName() string {
	return a.person.first + " " + a.person.last + " " + a.penName
}
// define author and embed person
type author struct {
	person
	penName string
	booksPublished int
}
//

// override fullName method for author
//

func main() {
	// initialize and print a person's full name
	p := person{
		first: "Toni",
		last:  "Morrison",
	}
	fmt.Println(p.fullName())

	// initialize and print an author's full name
	a := author {
		person: person{
			first : "JK",
			last : "Rowling",
		},
		penName: "JK",
		booksPublished: 10,
	}
	//
	fmt.Println(a.fullName())
}
