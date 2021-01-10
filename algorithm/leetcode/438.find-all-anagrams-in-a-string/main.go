package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
)

var _ = []tags.Array{tags.SlidingWindow}

func findAnagrams(s string, p string) []int {
	var ans []int
	if len(s) < len(p) {
		return ans
	}
	standard := [26]int{}
	mapping := [26]int{}
	for _, r := range []rune(p) {
		standard[r-'a']++
	}
	runes := []rune(s)
	for l, r := 0, -1; l <= len(s)-len(p); {
		if r < len(runes) && r-l+1 < len(p) {
			r++
			mapping[runes[r]-'a']++
			continue
		}
		if standard == mapping {
			ans = append(ans, l)
		}
		mapping[runes[l]-'a']--
		l++
	}
	return ans
}

func main() {
	fmt.Printf("%v", findAnagrams("cbaebabacd", "abc"))
}
