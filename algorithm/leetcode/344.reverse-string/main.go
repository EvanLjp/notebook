package main

import "github.com/evanljp/algorithm/tags"

var _ = []tags.Array{tags.CollisionPointer}

func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func main() {

}
