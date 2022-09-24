package main
import "fmt"

func main(){
	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("the %d th x value is : %d\n", i, x)
	}
}