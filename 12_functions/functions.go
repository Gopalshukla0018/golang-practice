package main

import "fmt"

// function for add two numbers
func add(a int, b int) int {
	return a + b

}

// if parameter is same typ we can give tyop only 1 ime
// ex--

func sum(a, b int) int {
	return a + b

}

// functions in go can returns multiple values

func getLanguages() (string, string, string, string, int) {
	return "golang", "javascript", "c", "python", 10
}

// functions in go is first class citizen so we can  assign in to a variable and  pass insite a function as argument and also return new fn under a function

// func processIt(fn func(a int)int){
// 	fn(1)
// }
func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}
func main() {
	res := add(3, 5)
	// fn:= func(a int)int{
	// 	return 2
	// }
	fn := processIt()
	fn(6)

	fmt.Println(res)
	resSum := sum(10, 20)
	fmt.Println(resSum)
	fmt.Println(getLanguages()) // working fine but we can get value like--..
	// example-
	lang1, leng2, lang3, lang4, lang5 := getLanguages()

	fmt.Println(lang1, leng2, lang3, lang4, lang5)
}
