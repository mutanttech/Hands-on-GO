// types/variables/begin/main.go
package main

import (
	"fmt"
	"strconv"
)

// declare package-level variables of type int
var age, ssn int
//
// declare package-level variables of type string
var name, city, country string
//

// declare package-level variables of type bool and override the default values (also known as "zero")
var drinks, smokes, party = true, true, true
//
func printVars(){
	//fmt.Println(x, y, z)
	fmt.Println(age, ssn)
}

func main() {
	
	// print ints
	name = "David"
	city = "Amsterdam"
	country = "Netherlands"
	fmt.Println(name, city, country)
	age = 33
	ssn = 1234
	fmt.Println(strconv.Itoa(age), strconv.Itoa(ssn))
	fmt.Println(drinks, smokes, party)

	x,y,z := 1,1,1
	fmt.Println(x, y, z)
	printVars()
	//

	// override the default value of a package-level variable
	// d = 1_000_000
	// fmt.Printf("d: %d\n", d)

	// declare and initialize variables of similar names as booleans but of local scope
	// x, y, z := 0, 1.25, "hello"
	// fmt.Println("x, y, z:", x, y, z)

	// print the package-level variables x, y, and z
}
