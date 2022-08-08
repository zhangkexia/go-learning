package main

import "fmt"

func main() {
    /* 局部变量定义 */
    var a int = 100;

    /* 判断布尔表达式 */
    if a < 20 {
        fmt.Printf("a less than 20\n");
    } else {
        fmt.Printf("a not less than 20\n");
    }
    fmt.Printf("a is %d\n", a)
}