package main

import "fmt"

func main() {
    sum :=1
    for ; sum <= 10; {
        sum += sum
    }
    fmt.Println(sum)

    //  也可以这样写
    for sum <= 10 {
        sum += sum
    }

    fmt.Println(sum)
}