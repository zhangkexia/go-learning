package main
import "fmt"

var x, y int
var ( //这种因式分解关键字的写法，一般用于声明全局变量
    v int
    w bool
)
var cc, dd int = 1, 2
var ee, ff = 123, "hello"

//这种不带声明格式的只能在函数体中出现
// gg, hh := 123, "hello"

func main() {
    //var a string = "Runoob"
    //fmt.Println(a)

    //var b, c int = 1, 2
    //fmt.Println(b, c)

    // 声明一个变量并初始化
    var a = "RUNOOB"
    fmt.Println(a)

    // 没有初始化就是零值
    var b int
    fmt.Println(b)

    // bool 零值为false
    var c bool
    fmt.Println(c)

    var i int
    var f float64
    var g bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, g, s)

    var d = true
    fmt.Println(d)

    intVal := 1 // 相等于 var intVal int; intVal = 1
    fmt.Println(intVal)

    k := "Runoob" // var f sting = "Runoob"

    fmt.Println(k)

    gg, hh := 123, "hello"
    println(x, y, a, b, cc, dd, ee, ff, gg, hh)
}