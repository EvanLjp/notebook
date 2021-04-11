package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
	"sort"
)

var _ = []tags.LUT{tags.LookupTable}

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		target := -num
		for start, end := i+1, len(nums)-1; start < end; {
			if start > i+1 && nums[start] == nums[start-1] {
				start++
				continue
			}
			switch {
			case nums[end]+nums[start] == target:
				res = append(res, []int{nums[i], nums[start], nums[end]})
				start++
			case nums[end]+nums[start] > target:
				end--
			case nums[end]+nums[start] < target:
				start++
			}
		}
	}
	return res
}

func main() {
	fmt.Printf("%v\n", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Printf("%v\n", threeSum([]int{0, 0, 0, 0, 0, 0}))
	fmt.Printf("%v\n", threeSum([]int{1, 1, -2}))
}
