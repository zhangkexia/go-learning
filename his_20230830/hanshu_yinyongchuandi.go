package main    

import "fmt"

func main() {
    /* 定义局部变量 */
    var a int = 100
    var b int = 200

    fmt.Printf("交换前， a的值：%d\n", a)
    fmt.Printf("交换前， b的值：%d\n", b)

    /* 调用 swap() 函数
    * &a 指向a 指针， a变量的地址
    * &b 指向b 指针， b变量的地址
    */
    swap(&a, &b)

    fmt.Printf("交换后， a的值： %d\n", a)
    fmt.Printf("交换后， b的值： %d\n", b)
}
 

// 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数
// 引用传递指针参数到函数内

/* 定义交换值函数 */
func swap(x *int, y *int) {
    var temp int
    temp = *x /* 保存x地址上的值 */
    *x = *y /* 将y值赋给x */
    *y = temp /* 将temp值赋给 y*/
}