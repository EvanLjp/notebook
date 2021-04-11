package main

import "github.com/evanljp/algorithm/tags"

var _ = []tags.LUT{tags.LookupTable}

func isAnagram(s string, t string) bool {
	chars := make([]byte, 26)
	for _, b := range []byte(s) {
		chars[b-'a']++
	}
	for _, b := range []byte(t) {
		chars[b-'a']--
	}
	for _, cnt := range chars {
		if cnt != 0 {
			return false
		}
	}
	return true
}

func main() {

}
