package main

import "fmt"

func main(){
	// boolean
	var answer bool //when undeclared is false
	fmt.Println(answer)

	// Numeric
	// int: int8 int16 int32 int64, uint8 uint16 uint32 uint64
	// int32 + int64 is error
	var a int = 10
	var b int = 3
	fmt.Println(a / b) // 10 // 3 not float 3.33333..
	// bit operation
	fmt.Println(a & b) // and
	fmt.Println(a | b) // or
	fmt.Println(a ^ b) // XOR
	fmt.Println(a &^ b) //

	// bit shifting
	fmt.Println(a << 3) // a * 2^3
	fmt.Println(a >> 3) // a / 2^3

	// floating point: float32, float64 default is float64
	var c float32 = 10.6
	var d float32 = 3.3
	fmt.Println(c + d)
	fmt.Println(c - d)
	fmt.Println(c * d)
	fmt.Println(c / d)

	// complex number: complex64, complex128
	var z1 complex64 = 1 + 2i
	var z2 complex64 = 3 + 4i
	var z3 complex128 = complex(5, 9)
	fmt.Printf("%T, %T\n", z1, z2)
	fmt.Printf("%v, %T\n", real(z1), real(z1))
	fmt.Printf("%v, %T\n", imag(z3), imag(z3))
	fmt.Println(z1 + z2)
	fmt.Println(z1 - z2)
	fmt.Println(z1 * z2)
	fmt.Println(z1 / z2)

	// strings utf8: 8bit
	var s1 string = "Hello World!"
	var s2 string = "From Go!"
	fmt.Println(s1 + s2)
	fmt.Printf("%v, %T\n", s1[2], s1[2]) // s1[2] is a char so %v returns the ascii value and therefore %T returns uint8
	fmt.Printf("%v, %T\n", string(s1[2]), string(s1[2])) // explicit conversion

	byteArray := []byte(s1) // an array where each index is a ascii to the char of s1
	fmt.Printf("%v, %T\n", byteArray, byteArray)

	// rune utf32 type: 32 bit
	var rn rune = 'a'
	fmt.Printf("%v, %T\n", rn, rn)

}
