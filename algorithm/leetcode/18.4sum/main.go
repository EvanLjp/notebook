package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
	"sort"
)

var _ = []tags.LUT{tags.LookupTable}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			start, end := j+1, len(nums)-1
			for start < end {
				if start > j+1 && nums[start] == nums[start-1] {
					start++
					continue
				}
				sum := nums[start] + nums[end] + nums[i] + nums[j]
				switch {
				case sum == target:
					res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})
					start++
				case sum < target:
					start++
				case sum > target:
					end--

				}
			}

		}

	}
	return res
}

func main() {
	fmt.Printf("%v", fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0))
}
