package main

func main() {
	canConstruct("a", "b")    // false
	canConstruct("aa", "ab")  // false
	canConstruct("aa", "aab") // true
}

func canConstruct(ransomNote string, magazine string) bool {
	ransomSlice := []rune(ransomNote)
	magazineSlice := []rune(magazine)

	for _, r := range ransomSlice {
		found := false
		for i, m := range magazineSlice {
			if r == m {
				magazineSlice = append(magazineSlice[:i], magazineSlice[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			//fmt.Println("false")
			return false
		}
	}
	//fmt.Println("true")
	return true
}
