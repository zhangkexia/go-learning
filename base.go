package main

import (
    "fmt"
)

func main() {
    // 格式化字符串
    var stock=123
    var enddate="2020-12-31"
    var url="Code=%d&endDate=%s"
    var target_url=fmt.Sprintf(url,stock,enddate) // fmt.Sprintf() 格式化字符串
    fmt.Println(target_url)

}