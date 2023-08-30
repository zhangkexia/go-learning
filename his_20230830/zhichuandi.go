package main
import "fmt"

func main() {
	/* local vars */
	var a int = 100
	var b int = 200

	fmt.Printf("before swap-a: %d\n", a)``
	fmt.Printf("before swap-b: %d\n", b)

	/* call function */
	swap(a, b)

	fmt.Printf("after swap-a: %d\n", a)
	fmt.Printf("after swap-b: %d\n", b)
}

/* define function-swap */
func swap(x, y int) int {
	var temp int 

	temp = x 
	x = y 
	y = temp  

	return temp;
}