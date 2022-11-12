package main

import "fmt"

/* statement global var */
var g int 

func main() {
    /* statement local var */
    var a, b int

    /* init var value */
    a = 10 
    b = 20
    g = a + b

    fmt.Printf("result: a = %d, b = %d and g = %d\n", a, b, g)
}