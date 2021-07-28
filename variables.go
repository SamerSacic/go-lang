/*
	- Variable declaration
	- Redeclaration and shadowing
	- Visibility
	- Naming convetions
	- Type conversions
*/
package main

import (
	"fmt"
	"strconv"
)

// declare variable in package level
// must be declared with type
var k float64 = 56

// redeclareVariable()
var i int = 27

// namingVariables()
var l int = 88
var L int = 90

// declare variables in block
var (
	actorName string = "Samer Sacic"
	companion string = "Sanel Sacic"
	doctorNumber int = 3
	season int = 11
)

func main() {
	declareVariable()
	redeclareVariable()
	namingVariables()
	variableConversion()
	fmt.Println(actorName, companion, doctorNumber, season)
}

func declareVariable() {
	// three ways to declare variables
	var a int 
	a = 42

	var b float32 = 16

	c := 65

	fmt.Println(a, b, c, k)
	fmt.Printf("%v, %T", b, b)
}

func redeclareVariable() {
	// variables in go must be used 
	// you cannot have variable that is not used
	fmt.Println(i)
	// variable with inner scope have precedence 
	// in other word that is shadowing
	// variable i is declared in package level and hare again
	var i int = 42
	// cannot redeclare a new variable
	// i := 13
	
	// can assign new value to variable
	// i = 13
	fmt.Println(i)
}

func namingVariables() {
	// lowercase variables defined on package level is used just in package
	fmt.Println(l)

	// uppercase variables is exposed for outside world
	fmt.Println(L)

	// we have three levels variable visibility in go
	// 1. Package level variable lowercase is scoped to the package (any file in the same package can acesse to the variable)
	// 2. Package level variable uppercase is exposed for outside of package, globally visible 
	// 3. Block scope (inside the functions)

	// acronyms should be always UPPERCASE
	var theURL string = "https://google.com"
	fmt.Println(theURL)
}

func variableConversion() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i) // conversion function or conversion operator 
	fmt.Printf("%v, %T\n", j, j)

	// NOTE: when convert from float to int need to be careful because we lose some decimal points
	// NOTE: you must make an explicit conversion, go does not support implicit conversions with variables

	var number int = 42
	fmt.Printf("%v, %T\n", number, number)

	// conversion from int to string yields a string of one rune, not a string of digits
	// var text string = string(number)
	var text string = strconv.Itoa(number)
	fmt.Printf("%v, %T\n", text, text)
 }

 /*
 	SUMMARY

	Variable declaration:
 		- var foo int
		- var foo int = 42
		- foo := 42

	Can't redeclare variables, but can shadow them
 	All variables must be used
	Visibility
 		- lower case first letter for package scope
		- upper case first letter to export
		- no private scope
	Naming conventions
 		- Pascal or camelCase
 			- Capitalize acronyms (HTTP, URL etc)
		- As short as reasonable
 			- longer names for longer lives
	Type conversions
 		- destinationType(variable)
		- use strconv package for strings
		- go don't have implicit conversions
 */