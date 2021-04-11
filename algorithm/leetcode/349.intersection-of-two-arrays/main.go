package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
)

var _ = []tags.LUT{tags.LookupTable}

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	for _, val := range nums1 {
		m[val] = struct{}{}
	}
	set := make(map[int]struct{})
	for _, val := range nums2 {
		if _, ok := m[val]; ok {
			set[val] = struct{}{}
		}
	}
	var res []int
	for i := range set {
		res = append(res, i)
	}
	return res
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	fmt.Printf("%v", intersection(nums1, nums2))
}
