package main

import "fmt"

func main() {
    // define local var
    var a int = 100
    var b int = 200

    fmt.Printf("before swap, a is %d\n", a)
    fmt.Printf("before swap, b is %d\n", b)

    // 调用函数用于交换值 &a指向a的地址；&b指向b的地址

    swap(&a, &b);

    fmt.Printf("after swap, a is %d\n", a)
    fmt.Printf("after swap, b is %d\n", b)
}

func swap(x *int, y *int) {
    var temp int
    temp = *x
    *x = *y
    *y = temp
    // 使用指针交换值，函数运行是不会额外创建新的局部变量
    // 如果使用值传递，形式参数x, y会在函数运行时创建局部变量，并进行交换，不影响a, b的值

}