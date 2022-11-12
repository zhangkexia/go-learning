package main

import "fmt"

func main() {
    var i, j, k int
    // state array and init value
    balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

    // output elements 
    for i = 0; i < 5; i++ {
        fmt.Printf("balance[%d] = %f\n", i, balance[i])
    }

    balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    // output elements
    for j = 0; j < 5; j++ {
        fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
    }


    // init element of index 1 and index 3
    balance3 := [5]float32{1:2.0, 3:7.0}
    for k = 0; k < 5; k++ {
        fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
    }
}