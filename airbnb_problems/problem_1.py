"""This problem was asked by Airbnb.
Given a list of integers, write a function that returns the largest sum of
non-adjacent numbers. Numbers can be 0 or negative.
For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5.
[5, 1, 1, 5] should return 10, since we pick 5 and 5.
Follow-up: Can you do this in O(N) time and constant space?
"""
"""
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
func main() {
	// pick n-th element (n needs to be iniated at 0 or 1)
	// get skip the n+1th element
	// get n+2 -> recursively
	// get n+3 -> recursively

	// return compare max
	memo := map[int]int{}

	data := []int{1, 5, 1, 1, 5}

	fmt.Println(math.Max(float64(getMaxSum(data, 0, memo)), float64(getMaxSum(data, 1, memo))))

}
"""