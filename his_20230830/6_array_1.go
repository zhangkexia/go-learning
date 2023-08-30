package main

import "fmt"

func main() {

    var n [10]int /* n is an arrary, which length is 10*/
    var i, j int

    /* init elements for array n */
    for i = 0; i < 10; i++ {
        n[i] = i + 100 /* set element is i + 100 */
    }

    /* output value of every element */
    for j = 0; j < 10; j++ {
        fmt.Printf("Element[%d] = %d\n", j, n[j])
    }
}