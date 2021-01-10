package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
)

var _ = []tags.Array{tags.SlidingWindow}

func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	mapping := make([]int, 256)
	length := 0
	runes := []rune(s)
	for _, char := range runes {
		if mapping[char] > 0 {
			for mapping[char] > 0 {
				mapping[runes[l]]--
				l++
			}
		}
		mapping[char]++
		r++
		if r-l > length {
			length = r - l
		}
	}
	return length
}

func main() {
	fmt.Printf("%d", lengthOfLongestSubstring("aab"))
}
