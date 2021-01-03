package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
)

var _ = []tags.Array{tags.Scan}

func moveZeroes(nums []int) {
	p := 0
	for _, num := range nums {
		if num != 0 {
			nums[p] = num
			p++
		}
	}
	for i := p; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroes2(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != k {
				nums[k], nums[i] = nums[i], nums[k]
			}
			k++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes2(nums)
	fmt.Printf("%v", nums)
}
