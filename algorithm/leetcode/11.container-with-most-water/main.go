package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
)

var _ = []tags.Array{tags.CollisionPointer}

func maxArea(height []int) int {
	res := 0
	lastHeight := 0
	for l, r := 0, len(height)-1; l < r; {
		if height[l] <= lastHeight {
			l++
			continue
		} else if height[r] <= lastHeight {
			r--
			continue
		}
		val := area(height[l], height[r], l, r)
		if val > res {
			res = val
		}
		lastHeight = min(height[l], height[r])
	}
	return res
}

func area(a, b, l, r int) int {
	if a <= b {
		return (a) * (r - l)
	} else {
		return (b) * (r - l)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("%d", maxArea(arr))
}
