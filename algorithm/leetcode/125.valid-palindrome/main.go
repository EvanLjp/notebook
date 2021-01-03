package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
)

var _ = []tags.Array{tags.CollisionPointer}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if !isAlphanumeric(s[l]) {
			l++
			continue
		}
		if !isAlphanumeric(s[r]) {
			r--
			continue
		}
		if isEqual(s[l], s[r]) {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func isEqual(l uint8, r uint8) bool {
	if l == r {
		return true
	}
	if l >= 'A' && l <= 'Z' {
		return (l + 32) == r
	}

	if r >= 'A' && r <= 'Z' {
		return (r + 32) == l
	}
	return false
}

func isAlphanumeric(c uint8) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func main() {
	fmt.Printf("%v", isPalindrome("A man, a plan, a canal: Panama"))
}
