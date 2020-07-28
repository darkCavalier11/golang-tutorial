package main
import "fmt"

const (
	// iterators
	i0 = iota // 0
	i1 // 1
	i2 // 2
	i3 // 3
	i4 // 4
)

func main(){
	// only primitive type can be assigned to a const
	// a function statement cant be assigned as const
	// const sin float64 = math.Sin(1.57) is an error as Sin() is a function
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	fmt.Printf("%v\n", i0)
	fmt.Printf("%v\n", i1)
	fmt.Printf("%v\n", i2)
	fmt.Printf("%v\n", i3)

}