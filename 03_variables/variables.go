package main

import "fmt"

func main() {

	/// If we declare a variable in Go, we must use it; otherwise, we have to remove it.

	// var name string = "Gopal"

	var name = "Gopal" // it is  also correct go infer the type

	// var isAuthenticated bool = true

	var isAuthenticated = true



	// shorthand syntax---
	price := 100;
	fmt.Println(price)
		// but we have to define type if assign value later 
	var rate int
	rate=200
	fmt.Println(rate)


	fmt.Println("authenticated is: ", isAuthenticated)

	fmt.Println(name)
}
