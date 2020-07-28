package main

import "fmt"
import "strconv" //for converting int and float to ascii

// global variables with a group that can be accessed
var (
	name string = "sumit kumar pradhan"
	age  int    = 45
)

func main() {
	// Printf formats according to a format specifier and writes to standard output.
	// It returns the number of bytes written and any write error encountered
	fmt.Printf("%v, %T\n", name, name)
	var str1 string
	var str2 string
	var weight float32 = 63.4
	var rounded int
	str1 = strconv.Itoa(age) // convert age = 18 to "18"
	str2 = string(age)       // return ascii value at position age
	fmt.Println(str1 + str2)

	// rounded = weight // error as converting float to int may loss some values
	rounded = int(weight) //  explicit conversion works for converting float to int
	fmt.Println(rounded)
}
