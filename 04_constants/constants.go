package main

import "fmt"

// asign outeside the main func
// but  can not not assign using :=
// const age = 20

func main() {
	// can not assign again
	const name string = "Gopal"
	const age = 20
	fmt.Println(name,age)

	const (
		host = "localhost://"
		port =500
	)

	fmt.Println(host,port)
}
