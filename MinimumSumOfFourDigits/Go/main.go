package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
note sure if this is working correctly
*/
func main() {
	fmt.Println(minimumSum(2932)) // 52
	fmt.Println(minimumSum(4009)) // 13
	fmt.Println(minimumSum(1000)) // 1
	fmt.Println(minimumSum(9999)) // 198
	fmt.Println(minimumSum(1001)) // 2
	// fmt.Println(minimumSum(0123)) leading zero integer does not work in Go?
}

func minimumSum(num int) int {
	// Convert the number to a string because we need to split the number into two parts
	numStr := strconv.Itoa(num)

	// Create a slice to store the digits
	digits := make([]int, 4)

	for i, c := range numStr {
		digits[i], _ = strconv.Atoi(string(c))
	}

	sort.Ints(digits)

	new1 := digits[0]*10 + digits[2]
	new2 := digits[1]*10 + digits[3]

	return new1 + new2

}
