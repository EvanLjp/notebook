package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
	"math"
)

var _ = []tags.LUT{tags.LookupTable}

func isHappy(n int) bool {
	dict := make(map[int]struct{})
	for {
		if _, ok := dict[n]; ok {
			return false
		}
		numbers := splitNumber(n)
		res := 0
		for _, number := range numbers {
			res += int(math.Pow(float64(number), 2))
		}
		if res == 1 {
			return true
		} else {
			dict[n] = struct{}{}
			n = res
		}
	}
}

func splitNumber(n int) []int {
	var res []int
	for n != 0 {
		remainder := n % 10
		n /= 10
		res = append(res, remainder)
	}
	return res
}

func main() {
	fmt.Printf("%v", isHappy(7))
}
