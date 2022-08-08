package main

import "fmt"

func main() {
    /* 定义局部变量 */
    var a int = 10

    /* 使用if语句判断不二表达式 */
    if a < 20 {
        fmt.Printf("a less than 20\n")
    }
    fmt.Printf("a value is %d\n", a)
}
