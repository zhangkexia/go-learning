package main

import "fmt"

func main() {
    var a int = 20 // state fact var
    var ip *int // state pointer var 

    ip = &a // pointe var stores address of fact var

    fmt.Printf("address of a is %x\n", &a)

    //value of ip (address of a)
    fmt.Printf("value of pointer ip %x\n", ip)

    //get var value by pointer ip
    fmt.Printf("value of pointer ip -> a %d\n", *ip)
}