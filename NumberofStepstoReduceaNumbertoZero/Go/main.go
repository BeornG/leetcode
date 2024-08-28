package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14)) // 6
	/*
		7
		6
		3
		2
		1
		0
	*/
}

func numberOfSteps(num int) int {

	res := 0
	for num > 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num - 1
		}
		res++
	}
	return res

}
