package main

import "fmt"

func main() {
    map1 := make(map[int]float32)
    map1[1] = 1.0
    map1[2] = 2.0
    map1[3] = 3.0
    map1[4] = 4.0

    // key and value
    for key, value := range map1 {
        fmt.Printf("key is %d - Value is %f\n", key, value)
    }

    // key
    for key := range map1 {
        fmt.Printf("key is %d\n", key)
    }

    //value
    for _, value := range map1 {
        fmt.Printf("value is %f\n", value)
    }
}