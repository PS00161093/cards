/**
Problem: https://leetcode.com/problems/implement-strstr/
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(getIndex("", ""))
}

func getIndex(hayStack, needle string) int {
	return strings.Index(hayStack, needle)
}
