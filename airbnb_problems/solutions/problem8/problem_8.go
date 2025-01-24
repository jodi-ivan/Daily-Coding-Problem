package solutions

import (
	"log"
	"math"
	"slices"
)

var bestSolution []float64
var smallestPairDiff float64

func Sum(s []float64) (total float64) {
	for _, v := range s {
		total += v
	}

	return total
}

func SumCeil(s []float64) (total float64) {
	for _, v := range s {
		total += math.Ceil(v)
	}

	return total
}

func AbsPairDiff(inital []float64, whole []float64) (res float64) {
	for i, v := range inital {
		res += math.Abs(v - whole[i])
	}

	return res
}

func getSolution(initial, arr []float64, index int, target float64) {

	if index > len(arr)-1 {
		return
	}

	if arr[index] == float64(target) {
		// sucess solution

		currAbsPairDiff := AbsPairDiff(initial, arr)
		log.Println("Found solution: ", arr, "with abs diff ", currAbsPairDiff)
		if smallestPairDiff == 0 {
			smallestPairDiff = currAbsPairDiff
			bestSolution = slices.Clone(arr)
		} else {
			if currAbsPairDiff < smallestPairDiff {
				smallestPairDiff = currAbsPairDiff
				bestSolution = slices.Clone(arr)
			}
		}
		return
	}

	if arr[index] > float64(target) {
		// not found solution
		return
	}

	target = target - arr[index]

	if smallestPairDiff != 0 && smallestPairDiff > AbsPairDiff(initial[0:index], arr[0:index]) {
		log.Println("Prunning at smallest pair diff ", arr)
	}

	if index < len(arr)-1 {

		if math.Round(Sum(arr[index+1:])) > target {
			log.Println("Prunning bc exceeds target: ", target, arr)
			return
		}
		ceil := slices.Clone(arr) // golang array is call by reference
		ceil[index+1] = math.Ceil(ceil[index+1])
		getSolution(initial, ceil, index+1, target)

		floor := slices.Clone(arr)
		floor[index+1] = math.Floor(floor[index+1])
		getSolution(initial, floor, index+1, target)

	}

}
func Solution8() {

	input := []float64{1.3, 2.3, 4.4, 7.1, 6.3, 6.9, 9.1, 2.3, 7.7}
	totalInput := math.Round(Sum(input))

	getSolution(input, append([]float64{math.Ceil(input[0])}, input[1:]...), 0, totalInput)
	getSolution(input, append([]float64{math.Floor(input[0])}, input[1:]...), 0, totalInput)

	log.Println("Best solution:", smallestPairDiff, "Array: ", bestSolution)
}
