package pkg

import (
	"fmt"
)

func recursiveSearching(n int64) int64 {
	var result int64
	if n == 0 {
		return result
	} else {
		if n > 0 {
			result = n * recursiveSearching(n-1)
		} else {
			result = n * recursiveSearching(n+1)
		}
	}
	return -1
}
func Calculate(n int64, check bool) int64 {
	if n == 0 {
		return 0
	}
	if check == true {
		fmt.Println("Start calculations...\nCalculate <", n, ">!")
	}
	result := recursiveSearching(n)
	fmt.Println("Calculations complete!")
	return result
}
