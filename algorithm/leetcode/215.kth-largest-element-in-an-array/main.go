package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
	"math/rand"
)

var _ = []tags.Array{tags.Partition}

func findKthLargest(nums []int, k int) int {
	m := len(nums) - k
	return findKth(nums, m, 0, len(nums)-1)
}

func findKth(nums []int, pos int, low, high int) int {
	for {
		left, right := partition(nums, low, high)
		if pos >= left && pos <= right {
			return nums[left]
		} else if pos < left {
			return findKth(nums, pos, low, left-1)
		} else {
			return findKth(nums, pos, right+1, high)
		}
	}
}

func partition(nums []int, low, high int) (left int, right int) {
	if low == high {
		return low, low
	}
	val := nums[low+rand.Intn(high-low)]
	l, r := low-1, high+1
	for i := l + 1; i < r; {
		switch {
		case nums[i] == val:
			i++
		case nums[i] > val:
			r--
			nums[r], nums[i] = nums[i], nums[r]
		case nums[i] < val:
			l++
			nums[l], nums[i] = nums[i], nums[l]
			i++
		}
	}
	return l + 1, r - 1
}

func main() {
	nums := []int{1}
	fmt.Printf("%v", findKthLargest(nums, 1))
}
