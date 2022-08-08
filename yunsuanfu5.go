package main

import "fmt"

func main() {
    var a int = 4
    var b int32
    var c float32
    var ptr *int

    /* 运算符实例 */
    fmt.Printf("line 1 - a type is %T\n", a);
    fmt.Printf("line 2 - b type is %T\n", b);
    fmt.Printf("line 3 - c type is %T\n", c);

    /* & and * */
    ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
    fmt.Printf("a value is %d\n", a);
    fmt.Printf("*ptr is %d\n", *ptr);


}