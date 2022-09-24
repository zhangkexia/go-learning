package main

import (
	"fmt"
	"math"
)

func main() {
	/* 声明函数变量 */
	getSqureRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSqureRoot(9))

}