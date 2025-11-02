package main

import "fmt"

// we can pass n no. of parameters in veriadic fn() 	fmt.Println(1,2,3,4,"hello")

// we can pass int only
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func main() {

	// result := sum(10, 20, 30)

	nums := []int{3, 4, 5, 6}
	result := sum(nums...)
	fmt.Println(result)

}
