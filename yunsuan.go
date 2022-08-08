package main

import "fmt"

func main() {
    var a int = 21
    var b int = 10
    var c int

    c = a + b
    fmt.Printf("line 1 - c 的值 %d\n", c)
    c = a - b
    fmt.Printf("line 2 - c value is %d\n", c)
    c = a * b 
    fmt.Printf("line 3 - c value is %d\n", c)
    c = a / b
    fmt.Printf("line 4 - c value is %d\n", c)
    c = a % b 
    fmt.Printf("line 5 - c value is %d\n", c)
    a++
    fmt.Printf("line 6 - a value is %d\n", a)
    a = 21 //重新赋值
    a -- 
    fmt.Printf("line 7 - a value is %d\n", a)
}