package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
)

var _ = []tags.Array{tags.DoublePointer, tags.Partition}

func sortColors(nums []int) {
	// [0,l] means 0, [r,len(nums)-1] means 2
	// initial ranges: [0,-1] [len(nums),len(nums)-1]
	l, r := -1, len(nums)
	for i := 0; i < r; {
		switch nums[i] {
		case 1:
			i++
		case 0:
			l++
			nums[l], nums[i] = nums[i], nums[l]
			i++
		case 2:
			r--
			nums[i], nums[r] = nums[r], nums[i]
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Printf("%v", nums)
}
