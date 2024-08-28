package main

import "fmt"

func main() {

	fmt.Println(runningSum([]int{1, 2, 3, 4})) // [1, 3, 6, 10]

}

func runningSum(nums []int) []int {

	res := make([]int, len(nums))

	sum := 0

	for i := range nums {
		sum += nums[i]
		res[i] = sum
	}

	return res

}
