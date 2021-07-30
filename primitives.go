/*
	- Boolean type
	- Numeric type
		- Integers
		- Floating point
		- Complex numbers
	- Text types
*/
package main

import (
	"fmt"
)

func primitives() {
	fmt.Println("==============PRIMITIVES==============")
	booleanTypes()
	numericTypes()
	textTypes()
}

func booleanTypes() {
	// coomon case for using boolean is for creating
	// state flags in application
	// other coomon case for using boolean is for 
	// logical tests
	// n := 1 == 1 // true
	// m := 1 == 2 // false
	var n bool = true;
	fmt.Printf("%v, %T\n", n, n)

	// in go when initialize variable with primitives
	// the value of variable is always 0 in boolean case is false
	var m bool
	fmt.Printf("%v, %T\n", m, m)
}

func numericTypes() {
	integerTypes()
	floatingPointTypes()
	complexTypes()
}

func integerTypes() {
	// integer
	a := 42 // int unspecific size
	fmt.Printf("%v, %T\n", a, a)

	// integer types with specific size
	var (				// sizes
		b int8 = 1		// -128 - 127
		c int16 = 2		// -32 768 - 32 767
		d int32 = 3 	// -2 147 483 648 - 2 147 483 647
		e int64 = 4		// -9 223 372 036 854 775 808 - 9 223 372 036 854 775 807 
	)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)

	// unsigned integer
	var f uint = 42
	fmt.Printf("%v, %T\n", f, f)

	// unsigned integer types with specific size
	var (     
		g uint8 = 5		// 0 - 255
		h uint16 = 6	// 0 - 65 535
		j uint32 = 7	// 0 - 4 294 697 295
	)
	fmt.Printf("%v, %T\n", g, g)
	fmt.Printf("%v, %T\n", h, h)
	fmt.Printf("%v, %T\n", j, j)

	// arithmetic operations on integers
	x := 10
	y := 3
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y) // reminder operator only avaliable for integers

	// bit operators - avaliable just for integers
	fmt.Println(x & y) 	// 0010 = 2
	fmt.Println(x | y)	// 1011 = 11
	fmt.Println(a ^ y)	// 1001 = 9
	fmt.Println(x &^ y)	// 0100 = 8

	// bitshift operators - avaliable just for integers
	x = 8 // 2^3
	fmt.Println(x << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(x >> 3) // 2^3 / 2^3 = 2^0 = 1
}

func floatingPointTypes() {
	// floating point
	// float32 ±1.18E-38 - ±3.4E38
	// float64 ±2.23E-208 - ±1.8E308

	// defining floating numbers
	n := 3.14
	n = 13.7e72 // exponential 
	n = 2.1E14
	fmt.Printf("%v, %T\n", n, n)

	// explicit definition
	var a float32 = 3.14
	var b float64 = 8.10
	fmt.Printf("%v, %T\n", a, b)
	fmt.Printf("%v, %T\n", b, b)

	// arithmetic operations on float point numbers
	x := 10.2
	y := 3.7
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
}

func complexTypes() {
	// complex numbers
	// complex64
	// complex 128

	// defining complex numbers
	var c complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", c, c)

	// arithmetic operations on complex numbers
	x := 1 + 2i
	y := 2 + 5.2i
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)

	// decompose complex numbers
	var n complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", real(n), real(n))
	fmt.Printf("%v, %T\n", imag(n), imag(n))

	// creating complex numbers
	var a complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", a, a)
}

func textTypes() {
	// string
	// string represent UTF-8 character
	// string in go are aliases for bytes
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	// we can threat string like array collection of symbols
	fmt.Printf("%v, %T\n", s[2], s[2]) 
	
	// string concatenation
	s1 := "this is string 1"
	s2 := "this is string 2"
	fmt.Printf("%v, %T\n", s1 + s2, s1 + s2)

	// slice of bytes for string
	// uint8 is type alias for bytes
	// lot of functions in go are working with slice of bytes
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
	
	// rune
	// rune represent UTF-32 character
	// rune is the same of int32
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)
}

/*
	SUMMARY

	- Boolean type
		- Values are TRUE and FALSE
		- Not an alias for other type (e.g int)
		- Zero value is FALSE
	
	- Numeric type
		- Integers
			- Signed integers
				- int type has varying size, but min 32 bits
				- 8 bit (int8) trough 64 bit (int64)
			- Unsigned integers
				- 8 bit (byte and uint8) trough 32 bit (uint32)
			- Arithmetic operatios
			 	- Addition, subtraction, multiplication, division, remainder 
			- Bitwise operations
				- And, or, xor, and not
			- Zero value is 0
			- Can't mix types in same family (uint16 + uint32 = error)
		
		- Floating point numbers
			- Follow IEEE-754 standard
			- Zero value is 0
			- 32 and 64 bit versions
			- Literal style
				- Decimal (3.14)
				- Exponential (13e18 or 2E10)
				- Mixed (13.7e12)
			- Arithmetic operatios
			 	- Addition, subtraction, multiplication, division
				 
		- Complex numbers
			- Zero value is (0 + 0i)
			- 64 and 128 bit versions
			- Built-in functions
				- complex - make complex number from two floats
				- real - get real part as float
				- imag - get imaginary part as float
			- Arithmetic operatios
			 	- Addition, subtraction, multiplication, division
	
	- Text types
		- Strings
			- UTF-8
			- Immutable
			- Can be concatedated with plus (+) operator
			- Can be converted to []byte
		- Rune
			- UTF-32
			- Alias for int32
			- Special methods normally required to process
				- e.g string.Reader#ReadRune
*/