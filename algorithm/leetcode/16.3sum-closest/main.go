package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
	"math"
	"sort"
)

var _ = []tags.LUT{tags.LookupTable}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	min := math.MaxInt64
	var res int
	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}
		for start, end := i+1, len(nums)-1; start < end; {
			sum := num + nums[end] + nums[start]
			if start > i+1 && nums[start] == nums[start-1] {
				start++
				continue
			}
			switch {
			case sum == target:
				return target
			case sum > target:
				if sum-target < min {
					min = sum - target
					res = sum
				}
				end--
			case sum < target:
				if target-sum < min {
					res = sum
					min = target - sum
				}
				start++
			}
		}
	}
	return res
}

func main() {
	fmt.Printf("%v", threeSumClosest([]int{1, 1, 1, 1}, 3))
}
