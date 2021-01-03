package main

import "github.com/evanljp/algorithm/tags"

var _ = []tags.Array{tags.Scan}

func removeDuplicates(nums []int) int {
	p := 0
	for i := range nums {
		if nums[i] != nums[p] {
			p++
			nums[p] = nums[i]
		}
	}
	return p + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	l := removeDuplicates(nums)
	for i := 0; i < l; i++ {
		println(nums[i])
	}
}
