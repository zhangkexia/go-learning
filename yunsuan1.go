package main

import "fmt"

func main() {
    var a int = 21
    var b int = 10

    if( a == b ) {
        fmt.Printf("line 1 - a equal b\n")
    } else {
        fmt.Printf("line 1 - a not equal b\n")
    }

    if ( a < b ) {
        fmt.Printf("line 2 - a less than b\n")
    } else {
        fmt.Printf("line 2 - a not less than b\n")
    }

    if ( a > b ) {
        fmt.Printf("line 3 - a more than b\n")
    } else {
        fmt.Printf("line 3 - a not more than b\n")
    }
    /* Lets change value of a and b */

    a = 5
    b = 20

    if ( a <= b ) {
        fmt.Printf("line 4 - a less than or equal b\n")
    } else {
        fmt.Printf("line 4 - a not less than or equal b\n")
    }

    if ( b >= a ) {
        fmt.Printf("line 5 - b more than or equal b\n")
    }
}