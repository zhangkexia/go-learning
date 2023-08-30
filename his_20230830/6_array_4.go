package main

import "fmt"

func main() {
    // create 2-D array
    sites := [2][2]string{}

    // append elements 
    sites[0][0] = "Google"
    sites[0][1] = "Runoob"
    sites[1][0] = "Taobao"
    sites[1][1] = "Weibo"

    // show resutls
    fmt.Println(sites)
}