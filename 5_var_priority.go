package main

import "fmt"

/* state global var */
var g int = 20


/* priority of local var in function is higher than the same name global var */
func main() {
    /* state local var */
    var g int = 10

    fmt.Printf("result: g = %d\n", g)
}