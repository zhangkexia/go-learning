package main

import "fmt"

// go 语言支持匿名函数，可作为闭包。 匿名函数是一个“内联”语句表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
// 以下实例，创建函数getSquence(), 返回另外一个函数。 该函数目的时在闭包中递增i变量

func getSequence() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main() {
    /* nextNumber 为一个函数，函数i为0 */
    nextNumber := getSequence()

    /* 调用nextNumber函数， i 变量自增1并返回 */
    fmt.Println(nextNumber())
    fmt.Println(nextNumber())
    fmt.Println(nextNumber())

    /* 创建新的函数nextNumber1, 并查看结果 */
    nextNumber1 := getSequence()
    fmt.Println(nextNumber1())
    fmt.Println(nextNumber1())

}