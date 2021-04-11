package main

import "github.com/evanljp/algorithm/tags"

var _ = []tags.LUT{tags.LookupTable}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	dict := make(map[int]int)
	var count int
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			dict[num1+num2]++
		}
	}

	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			count += dict[-(num3 + num4)]
		}

	}
	return count

}

func main() {

}
