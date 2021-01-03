package main

import "github.com/evanljp/algorithm/tags"

var _ = []tags.Array{tags.Scan}

func removeElement(nums []int, val int) int {
	p := 0
	for i := range nums {
		if nums[i] != val {
			if i != p {
				nums[p], nums[i] = nums[i], nums[p]
			}
			p++
		}
	}
	return p
}

func main() {
	nums := []int{3, 2, 2, 3}
	l := removeElement(nums, 3)
	for i := 0; i < l; i++ {
		println(nums[i])
	}
}
