package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
	"math"
)

var _ = []tags.Array{tags.SlidingWindow}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	// make a mapping to count how many characters have.
	mapping := [123]int{}
	for _, v := range t {
		mapping[v]++
	}
	// the sliding window is [left,right]
	left, right, start, counter, minLen := 0, -1, 0, 0, math.MaxInt32
	for right < len(s)-1 {
		right++
		if mapping[s[right]] > 0 {
			counter++
		}
		mapping[s[right]]--
		for counter == len(t) {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}
			mapping[s[left]]++
			if mapping[s[left]] > 0 {
				counter--
			}
			left++
		}
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : start+minLen]
}

func main() {
	fmt.Printf("%s", minWindow("ADOBECODEBANC", "ABC"))
}
