package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))     // 3
	fmt.Println(romanToInt("IV"))      // 4
	fmt.Println(romanToInt("IX"))      // 9
	fmt.Println(romanToInt("LVIII"))   // 58
	fmt.Println(romanToInt("MCMXCIV")) // 1994
}

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	result := 0
	prev := 0
	// iterate the string from right to left to make it easier to compare the current value with the previous value
	for i := len(s) - 1; i >= 0; i-- {
		value := romanMap[string(s[i])]
		if value < prev { // if the value is less than the previous value, subtract it
			result -= value
		} else {
			result += value
		}
		prev = value
	}

	return result
}
