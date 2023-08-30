package main

import "fmt"

/* 声明全局变量 */
var g int = 20

func main() {
	/* 声明局部变量 */
	var g int = 10

	/* 同名变量情况下， 局部变量优先于全局变量 */

	fmt.Printf("result: g = %d \n", g)
}