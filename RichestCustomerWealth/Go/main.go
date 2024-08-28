package main

import "fmt"

func main() {

	fmt.Println(findRichestCustomer([][]int{{1, 2, 3}, {3, 2, 1}}))   // 6
	fmt.Println(findRichestCustomer([][]int{{1, 5}, {7, 3}, {3, 5}})) // 10
}

func findRichestCustomer(accounts [][]int) int {
	maxWealth := 0

	// iterate through each customer
	for _, customer := range accounts {
		wealth := runningSum(customer)[len(customer)-1] // reusing the runningSum function from RunningSumOf1dArray
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}

	return maxWealth
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
