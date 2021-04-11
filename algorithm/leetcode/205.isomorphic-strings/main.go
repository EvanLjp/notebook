package main

import "github.com/evanljp/algorithm/tags"

var _ = []tags.LUT{tags.LookupTable}

func isIsomorphic(s string, t string) bool {
	sarr := []rune(s)
	tarr := []rune(t)
	if len(sarr) != len(tarr) {
		return false
	}
	dict := make(map[rune]rune)
	used := make(map[rune]struct{})
	for i, sr := range sarr {
		val, ok := dict[sr]
		tr := tarr[i]
		if !ok {
			if _, ok = used[tr]; ok {
				return false
			} else {
				dict[sr] = tr
				used[tr] = struct{}{}
			}
		} else if val != tr {
			return false
		}
	}
	return true
}

func main() {

}
