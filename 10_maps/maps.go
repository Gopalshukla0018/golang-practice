package main

import (
	"fmt"
	"maps"
)

func main() {

	// creating map

	m := make(map[string]string)

	// setting an element
	m["firstName"] = "Gopal"
	m["lastName"] = "Shukla"

	fmt.Println(m["firstName"])
	fmt.Println(m["lastName"])
	//IMP if key does not exist in the map then it returns zero
	fmt.Println(m["name"]) // empty

	m2 := make(map[string]int)

	m2["age"] = 90
	m2["waight"] = 70
	fmt.Println(m2["age"])
	fmt.Println(m2["contact"]) // it willl give 0

	// length of map
	fmt.Println(len(m2)) // 2

	// delete element from map
	fmt.Println(m2) // it prints age and waight both
	delete(m2, "age")

	fmt.Println(m2) // it prints  on;y waight

	// clear map by clear() fn
	clear(m2)
	fmt.Println(m2) // empty map

	// create map without make() fn

	m3 := map[string]int{"price": 40, "agr": 20}
	fmt.Println(m3["price"]) // 40
	fmt.Println(m2)

	v, ok := m3["price"] // alll ok
	// _,ok :=m3["class"] not ok
	fmt.Println(v)

	if ok {
		fmt.Println("all ok")

	} else {
		fmt.Println("not ok")
	}

	m4 := map[string]int{"price": 40, "age": 20}
	m5 := map[string]int{"price": 40, "age": 20}

	//  fmt.Println(m4==m5) // we can not do this

	// there is a maps for this
	fmt.Println(maps.Equal(m4, m5)) // it returns true || false

}
