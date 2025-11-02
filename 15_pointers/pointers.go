package main

import "fmt"

// by value paas hua hai
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("int changeNum",num)
// }

// by reference

func changeNum(num *int) {
	*num = 5
	fmt.Println("in change num  ", *num)
}

func main() {
	num := 1
	changeNum(&num)

	fmt.Println("memory adress", &num) /// check memory adress
	fmt.Println("after changeNum in main", num)
}
