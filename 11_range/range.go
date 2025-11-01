package main

import "fmt"

// iterating over data structures
func main() {

	nums := []int{6, 7, 8}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }
	// sum := 0
	for i, num := range nums {

		// sum = sum + num
		fmt.Println(num, i)

	}
	// fmt.Println(sum)

	// we can iterate map by range
	m := map[string]string{"name": "Gopal", "surname": "Shukla"}

	for key, val := range m {
		fmt.Println(key, val)

	}
	// if u want only key so write only single value
	for k := range m {
		fmt.Println(k) // it prints only key

	}

	//unicod code point rune
	// starting byte of rune
	/// 255 x 300- -> byte,2byte
	for i, c := range "golang" {
		fmt.Println(i, c)
	}

}
