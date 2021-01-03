package main

import "github.com/evanljp/algorithm/tags"

var _ = []tags.Array{tags.CollisionPointer}

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for {
		val := numbers[l] + numbers[r]
		switch {
		case val == target:
			goto fin
		case val < target:
			l++
		case val > target:
			r--
		}
	}
fin:
	return []int{l + 1, r + 1}
}

func main() {

}
