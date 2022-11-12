package main

import "fmt"

func main() {
    // create empty 2-D array
    animals := [][]string{}

    // create 3 1-D array
    row1 := []string{"fish", "shark", "eel"}
    row2 := []string{"bird"}
    row3 := []string{"lizard", "salamander"}

    // append 1-D array to 2-D array
    animals = append(animals, row1)
    animals = append(animals, row2)
    animals = append(animals, row3)

    // outpu for loop
    for i := range animals {
        fmt.Printf("Row: %v\n", i)
        fmt.Println(animals[i])
    }
}