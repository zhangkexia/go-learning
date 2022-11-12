package main

import "fmt"

func main() {
    /* 定义局部变量 */
    var a int = 100
    var b int = 200

    fmt.Printf("交换前 a的值为： %d\n", a)
    fmt.Printf("交换前 b的值为： %d\n", b)

    /* 通过调用函数来交换值 */
    swap(a, b)

    fmt.Printf("交换后 a的值： %d\n", a)
    fmt.Printf("交换后 b的值： %d\n", b)

    // 使用的是值传递，值的互换仅在函数内部实现，不影响原有变量的值；
    // 值传递时，相当于复制了2个变量（作用域函数内部），新的2个变量完成交换，并不影响原有2个变量的值
}


/* 定义相互交换值的函数 */
func swap(x, y int) int {
    var temp int

    temp = x /* 保存x的值 */
    x = y /* 将y的值赋给 x */
    y = temp /* 将temp的值赋给y */

    return temp;
}
