/**
 * Problem URL: https://leetcode.com/problems/longest-palindromic-substring/
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getLongestPalindromeSubString("babad"))
}

func getLongestPalindromeSubString(str string) string {

	if len(str) < 1 {
		return ""
	}

	startIndex := 0
	endIndex := 0
	for i := 0; i < len(str); i++ {
		length := int(math.Max(expandAroundCenter(str, i, i), expandAroundCenter(str, i, i+1)))
		if length > endIndex-startIndex {
			startIndex = i - (length-1)/2
			endIndex = i + length/2
		}
	}

	return str[startIndex : endIndex+1]
}

func expandAroundCenter(str string, i int, j int) float64 {

	tempLeft := i
	tempRight := j
	for tempLeft >= 0 && tempRight < len(str) && str[tempLeft:tempLeft+1] == str[tempRight:tempRight+1] {
		tempLeft--
		tempRight++
	}

	return float64(tempRight - tempLeft - 1)
}
