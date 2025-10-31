package main

import (
	"fmt"
)

func main() {

	// normal switch -

	// i := 0

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")

	// // default
	// default:
	// 	fmt.Println("others")
	// }
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it is weekEnd")
	// default:
	// 	fmt.Println("weekday")
	// }

	// type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("it is integer")
		case string:
			fmt.Println("it is string")
		case bool:

			fmt.Println("it is boolean")
		default:
			fmt.Println("others")

		}

	}
	whoAmI(1.00)

}
