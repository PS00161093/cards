package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {

	maxSize := 0

	helper := make(map[string]int)
	for i := 0; i < len(s); i++ {
		ch := s[i : i+1]
		if _, ok := helper[ch]; ok {
			removeTillRepeatingCharacter(helper, ch)
		} else {
			helper[ch] = 1
		}
		if maxSize < len(helper) {
			maxSize = len(helper)
		}
	}

	return maxSize
}

func removeTillRepeatingCharacter(helper map[string]int, ch string) {

	for key, _ := range helper {
		if key == ch {
			delete(helper, ch)
		}
	}
	delete(helper, ch)
	helper[ch] = 1
}
