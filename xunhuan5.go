package main

import "fmt"

func main() {
    /* 定义全局变量 */
    var i, j int

    for i = 2; i < 100; i++ {
        for j = 2; j <= (i / j); j++ {
            if(i % j == 0) {
            break; //发现因子
            }
        }
        if( j > (i / j)) {
            fmt.Printf("%d is sushu\n", i);
        }
    }

}