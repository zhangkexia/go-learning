package main

import "fmt"

func main() {
    var a int = 21
    var c int 

    c = a 
    fmt.Printf("line 1 - c is %d\n", c)

    c += a
    fmt.Printf("line 2 - c is %d\n", c)

    c -= a
    fmt.Printf("line 3 - c is %d\n", c)

    c *= a
    fmt.Printf("line 4 - c is %d\n", c)

    c /= 1
    fmt.Printf("line 5 - c is %d\n", c)

    c = 200;

    c <<= 2
    fmt.Printf("line 6 - c is %d\n", c)

    c >>= 2
    fmt.Printf("line 7 - c is %d\n", c)

    c &= 2
    fmt.Printf("line 8 - c is %d\n", c)

    c ^= 2
    fmt.Printf("line 9 - c is %d\n", c)

    c |= 2
    fmt.Printf("line 10 - c is %d\n", c)
}