package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
)

var _ = []tags.LUT{tags.LookupTable}

func numberOfBoomerangs(points [][]int) int {
	res := 0
	for i := 0; i < len(points); i++ {
		record := make(map[int]int, len(points))
		for j := 0; j < len(points); j++ {
			if j != i {
				record[dis(points[i], points[j])]++
			}
		}
		for _, cnt := range record {
			res += cnt * (cnt - 1)
		}
	}

	return res
}

func dis(pa, pb []int) int {
	return (pa[0]-pb[0])*(pa[0]-pb[0]) + (pa[1]-pb[1])*(pa[1]-pb[1])
}

func main() {

	fmt.Printf("%v", numberOfBoomerangs([][]int{
		{0, 0},
		{1, 0},
		{2, 0},
	}))

}
