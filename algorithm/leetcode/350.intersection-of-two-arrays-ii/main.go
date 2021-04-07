package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
)

var _ = []tags.LUT{tags.Map}

func intersect(nums1 []int, nums2 []int) []int {
	size := len(nums1)
	if size < len(nums2) {
		size = len(nums2)
	}
	dict := make(map[int]int, size)

	for _, num := range nums1 {
		old, ok := dict[num]
		if ok {
			dict[num] = old + 1
		} else {
			dict[num] = 1
		}
	}

	var res []int

	for _, num := range nums2 {
		old, ok := dict[num]
		if ok && old > 0 {
			res = append(res, num)
			dict[num]--
		}
	}
	return res
}

func main() {

	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	fmt.Printf("%v", intersect(nums1, nums2))
}
