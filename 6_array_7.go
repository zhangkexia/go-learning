package main

import "fmt"

func main() {
    // array length is 5
    var balance = [5]int {1000, 2, 3, 17, 50}
    var avg float32

    // pass array to func
    avg = getAverage(balance, 5);

    // output avg
    fmt.Printf("average value is %f\n", avg)
}




func getAverage(arr [5]int, size int) float32 {
    var i,sum int
    var avg float32

    for i = 0; i < size; i++ {
        sum += arr[i]
    }

    avg = float32(sum) / float32(size)

    return avg;
}