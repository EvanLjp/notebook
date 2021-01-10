package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
)

var _ = []tags.Array{tags.SlidingWindow}

func minSubArrayLen(s int, nums []int) int {
	l, r := 0, 0
	sum := 0
	distance := 0
	for {
		if r == len(nums) && sum < s {
			break
		}
		if sum < s && r < len(nums) {
			sum += nums[r]
			r++
			continue
		}
		if sum >= s && l < len(nums) {
			if distance == 0 || distance > r-l {
				distance = r - l
			}
			sum -= nums[l]
			l++
			continue
		}
	}
	return distance
}

func main() {
	arr := []int{2, 3, 1, 2, 4, 3}
	distance := minSubArrayLen(7, arr)
	fmt.Printf("%d", distance)

}
