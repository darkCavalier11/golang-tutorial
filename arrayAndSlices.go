package main
import "fmt"
func main(){
	var student [3]string
	array := [3]int{4, 8, 6}
	fmt.Printf("%v, %T\n", array, array)
	student[0] = "sumit"
	fmt.Printf("%v, %T\n", student, student)
	fmt.Printf("%v\n", len(student))

	// 2D arrays
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	// copying of arrays and pointer
	a := [...]int{1, 2, 3}
	b := a
	c := &a
	b[1] = 5
	c[2] = 7
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// slices
	p := []int{1, 2, 3}
	fmt.Printf("Length %v\n", len(p))
	fmt.Printf("Capacity %v\n", cap(p))
	q := p // ref type. q store pointer
	q[0] = 5
	fmt.Printf("%v\n", p)

	// python type slicing
	r := p[:2] // r stored reference of p.
	fmt.Println(r)
	p[1] = 1 // changing p changes or not vice versa
	fmt.Println(p)
	fmt.Println(r)

	// slice are dynamic arrays
	
}