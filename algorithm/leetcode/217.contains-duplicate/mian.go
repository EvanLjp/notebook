package main

import "github.com/evanljp/algorithm/tags"

var _ = []tags.LUT{tags.LookupTable}

func containsDuplicate(nums []int) bool {
	dict := make(map[int]struct{})
	for _, num := range nums {
		_, ok := dict[num]
		if ok {
			return true
		}
		dict[num] = struct{}{}
	}
	return false
}

func main() {

}
