package main

import "fmt"

/* state global var */
var a int = 20;

func main() {
    /* state local var in main function */
    var a int = 10
    var b int = 20
    var c int = 0

    fmt.Printf("in main() function, a = %d\n", a);
    c = sum(a, b);
    fmt.Printf("in main() function, c = %d\n", c);
}

/* function defination-add two numbers */
func sum(a, b int) int {
    fmt.Printf("in sum() function, a = %d\n", a);
    fmt.Printf("in sum() function, b = %d\n", b);

    return a + b;
}