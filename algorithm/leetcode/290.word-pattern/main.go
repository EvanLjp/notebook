package main

import (
	"github.com/evanljp/algorithm/tags"
	"strings"
)

var _ = []tags.LUT{tags.LookupTable}

func wordPattern(pattern string, s string) bool {
	vals := strings.Split(s, " ")
	dict := make(map[rune]string)
	set := make(map[string]struct{})
	if len(vals) != len(pattern) {
		return false
	}
	for i, r := range pattern {
		m, ok := dict[r]
		_, has := set[vals[i]]
		if !ok && !has {
			dict[r] = vals[i]
			set[vals[i]] = struct{}{}
			continue
		}
		if m != vals[i] {
			return false
		}
	}
	return true
}

func main() {

}
