package main

import (
	"calculator/calc" // calc包所在的子目录路径，calc包的名字不一定要和子目录名称相同
	"fmt"
)

func main() {
	var m = 15
	var n = 20
	var r1 = calc.Add(m, n)
	var r2 = calc.Sub(m, n)
	fmt.Printf("%d + %d is %d\n", m, n, r1)
	fmt.Printf("%d - %d is %d\n", m, n, r2)
}
