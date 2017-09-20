package main

import (
	"fmt"
)

func removeDups(slice []int) []int {
	var finalSlice []int
	if len(slice) < 2 {
		return slice
	}

	finalSlice = append(finalSlice, slice[0])
	for i, _ := range slice {
		if i+1 == len(slice) {
			return finalSlice
		}
		if slice[i] != slice[i+1] {
			finalSlice = append(finalSlice, slice[i+1])
		}
	}
	return finalSlice
}

func main() {
	a := []int{0, 1}
	b := []int{1, 1}
	c := []int{0, 1, 0}
	d := []int{0, 1, 1}
	e := []int{1, 1, 0}
	f := []int{1, 1, 1}

	fmt.Println(removeDups(a)) // "[5 6 8 9]"
	fmt.Println(removeDups(b)) // "[5 6 8 9]"
	fmt.Println(removeDups(c)) // "[5 6 8 9]"
	fmt.Println(removeDups(d)) // "[5 6 8 9]"
	fmt.Println(removeDups(e)) // "[5 6 8 9]"
	fmt.Println(removeDups(f)) // "[5 6 8 9]"
}
