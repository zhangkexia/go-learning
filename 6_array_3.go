package main

import "fmt"

func main() {
    // step1: create array
    values := [][]int{}

    // step2: add one-dimensional array to two-dimentional array by append() function
    row1 := []int{1, 2, 3}
    row2 := []int{4, 5, 6}
    values = append(values, row1)
    values = append(values, row2)

    // step3: show 2 line data
    fmt.Println("Row 1")
    fmt.Println(values[0])
    fmt.Println("Row 2")
    fmt.Println(values[1])

    // step4: visit first elment
    fmt.Println("the first element is: ")
    fmt.Println(values[0][0])
}