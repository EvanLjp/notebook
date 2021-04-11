package main

import "github.com/evanljp/algorithm/tags"

var _ = []tags.LUT{tags.LookupTable}

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, num := range nums {
		index, ok := dict[target-num]
		if ok {
			return []int{index, i}
		}
		dict[num] = i
	}
	return []int{}
}

func main() {

}
