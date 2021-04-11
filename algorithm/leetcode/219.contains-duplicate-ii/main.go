package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
)

var _ = []tags.LUT{tags.LookupTable}

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]struct{})
	for i, num := range nums {
		_, ok := seen[num]
		if ok {
			return true
		}
		seen[num] = struct{}{}
		if i >= k {
			delete(seen, nums[i-k])
		}
	}
	return false
}

func main() {

	fmt.Printf("%v", containsNearbyDuplicate([]int{
		1, 2, 3, 1, 2, 3,
	}, 2))

}
