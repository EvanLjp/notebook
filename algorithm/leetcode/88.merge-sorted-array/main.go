package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
)

var _ = []tags.Array{tags.DoublePointer}

func merge(nums1 []int, m int, nums2 []int, n int) {
	last, l1, l2 := m+n-1, m-1, n-1
	for last > -1 && l1 > -1 && l2 > -1 {
		if nums1[l1] >= nums2[l2] {
			nums1[last] = nums1[l1]
			l1--
		} else {
			nums1[last] = nums2[l2]
			l2--
		}
		last--
	}
	for l1 > -1 {
		nums1[last] = nums1[l1]
		last--
		l1--
	}
	for l2 > -1 {
		nums1[last] = nums2[l2]
		last--
		l2--
	}
}

func main() {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	nums2 := []int{1, 2, 3}
	merge(nums1, 3, nums2, 3)
	fmt.Printf("%v", nums1)
}
