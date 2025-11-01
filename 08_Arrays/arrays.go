package main

import "fmt"

func main() {
	var nums [4]int

	   //len() for  get length of a array
	   fmt.Println(len(nums))

	nums[1] = 1
	// add 0 on rest..
	fmt.Println(nums)

	var vals [5]bool

	fmt.Println(vals) // by default all is false
	// by  default --- int-> 0.string-> "",bool -> false

	var name [3]string
	name[2]="Gopal"
	fmt.Println(name) 


	// to declare in single line
	newaArr := [3]int{1,2,3}
	fmt.Println(newaArr)


	// 2d array
	TwoDArr := [2][2]int {{1,2},{4,5}}
	fmt.Println(TwoDArr)



}
