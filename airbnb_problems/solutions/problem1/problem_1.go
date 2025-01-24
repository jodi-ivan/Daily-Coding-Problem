package solutions

import (
	"fmt"
	"math"
)

func getMaxSum(data []int, n int, memo map[int]int) int {
	if n > len(data)-1 {
		return 0
	}

	if val, exists := memo[n]; exists {
		return val
	}

	aSkip := getMaxSum(data, n+2, memo) + data[n]
	twoSkip := getMaxSum(data, n+3, memo) + data[n]

	return int(math.Max(float64(aSkip), float64(twoSkip)))
}
func Solution1() {
	// pick n-th element (n needs to be iniated at 0 or 1)
	// get skip the n+1th element
	// get n+2 -> recursively
	// get n+3 -> recursively

	// return compare max
	memo := map[int]int{}

	data := []int{1, 5, 1, 1, 5}

	fmt.Println(math.Max(float64(getMaxSum(data, 0, memo)), float64(getMaxSum(data, 1, memo))))

}
