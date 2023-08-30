package main

import "fmt"

func main() {
    /* 局部变量 */
    var grade string = "B"
    var marks int = 90

    switch marks {
        case 90: grade = "A"
        case 80: grade = "B"
        case 50, 60, 70 : grade = "C"
        default: grade = "D"
    }

    switch {
        case grade == "A":
            fmt.Printf("youxiu!\n")
        case grade == "B", grade == "C":
            fmt.Printf("lianghao!\n")
        case grade == "D":
            fmt.Printf("jige!\n")
        case grade == "F":
            fmt.Printf("bujige!\n")
        default:
            fmt.Printf("cha!\n");
    }
    fmt.Printf("nide dengji shi %s\n", grade);
}