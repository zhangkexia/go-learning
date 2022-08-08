package main

import "fmt"

func main() {
    var a bool = true
    var b bool = false
    if ( a && b ) {
        fmt.Printf("line 1, is true\n")
    }

    if ( a || b ) {
        fmt.Printf("line 2, is true\n")
    }

    /* modify value of a and b */
    a = false
    b = true
    if ( a && b ) {
        fmt.Printf("line 3, is true\n")
    } else {
        fmt.Printf("line 3, is true\n")
    }
    if ( !( a && b )) {
        fmt.Printf("line 4, is true\n")
    }
}