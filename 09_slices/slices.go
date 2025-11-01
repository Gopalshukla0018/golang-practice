package main

import "fmt"

// slices -> dynamic arrays
// flexible, powerful, and most used in Go
// unlike arrays, size can grow or shrink

func main() {
	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums)        // []
	fmt.Println(nums == nil) // true
	fmt.Println(len(nums))   // 0

	// if you don't want nil slice
	var sl = make([]int, 0, 7) // here 0 = length, 7 = capacity
	fmt.Println(sl)
	fmt.Println(sl == nil) // false
	fmt.Println(cap(sl))   // 7

	// adding values using append
	sl = append(sl, 1)
	sl = append(sl, 2)
	sl = append(sl, 3)
	sl = append(sl, 4)
	sl = append(sl, 5)
	sl = append(sl, 6)
	sl = append(sl, 7)
	sl = append(sl, 8)
	fmt.Println(sl)
	fmt.Println("capacity:", cap(sl)) // capacity automatically increases (doubles)

	// slicing -> taking a portion of slice
	part := sl[2:5] // index 2 se 4 tak (5 exclusive)
	fmt.Println("part:", part)

	// if you change the sub-slice, main slice also changes (same memory)
	part[0] = 999
	fmt.Println("after modifying part slice:")
	fmt.Println("part:", part)
	fmt.Println("sl:", sl)

	// copy() function
	src := []int{10, 20, 30, 40, 50}
	dst := make([]int, len(src))
	n := copy(dst, src)
	fmt.Println("Copied elements:", n)
	fmt.Println("dst slice:", dst)

	// append one slice into another
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	a = append(a, b...) // note the three dots (...)
	fmt.Println("appended slice:", a)

	// deleting element from slice
	// delete index 2 element (value 3)
	idx := 2
	a = append(a[:idx], a[idx+1:]...) // skip index 2
	fmt.Println("after delete:", a)

	// copy by reference concept
	s1 := []int{1, 2, 3}
	s2 := s1
	s2[0] = 100
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	// both change because slices hold reference to same array

	// make independent copy
	s3 := make([]int, len(s1))
	copy(s3, s1)
	s3[1] = 999
	fmt.Println("s1 (original):", s1)
	fmt.Println("s3 (independent copy):", s3)

	// length and capacity demo
	x := []int{1, 2, 3, 4, 5}
	fmt.Println("len:", len(x), "cap:", cap(x))
	y := x[:3]
	fmt.Println("y:", y, "len:", len(y), "cap:", cap(y)) // capacity same, just smaller view
}
