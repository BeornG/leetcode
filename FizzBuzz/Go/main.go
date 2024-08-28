package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(3))  // Output: ["1","2","Fizz"]
	fmt.Println(fizzBuzz(5))  // Output: ["1","2","Fizz","4","Buzz"]
	fmt.Println(fizzBuzz(15)) // Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
}

func fizzBuzz(n int) []string {
	array := []string{}

	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			array = append(array, "FizzBuzz")
		case i%3 == 0:
			array = append(array, "Fizz")
		case i%5 == 0:
			array = append(array, "Buzz")
		default:
			array = append(array, strconv.Itoa(i))
		}
	}

	return array
}
