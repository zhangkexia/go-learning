package main

import "fmt"

func main() {
	/* local vars */

	var a int = 100
	var b int = 200
	var ret int

	/* call function get max value */
	ret = max(a, b)

	fmt.Printf("max value is : %d\n", ret)
}

/* return max value of two numbers */
func max(num1, num2 int) int {
	/* local vars */
	var result int 

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}