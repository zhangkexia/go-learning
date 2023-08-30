package main

import "fmt"

func main() {
	// step1: 创建数组
	values := [][]int{}

	// step2: 使用append()函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	// step3: 显示两行数据
	fmt.Println("row 1")
	fmt.Println(values[0])
	fmt.Println("row 2")
	fmt.Println(values[1])

	// step4: 访问第一个元素
	fmt.Println("第一个元素为：")
	fmt.Println(values[0][0])
}