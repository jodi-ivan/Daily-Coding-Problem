package solutions

import (
	"log"
	"slices"
	"sort"
)

func Alt1() {
	/*
	 sort decending the array
	 for each element do
	   fill the array
	   if (curr_element - next_element == 1) then
	      fill the array
	   else
	      calculate sum of the array
	      max compare with the global
	      empty the array
	      if sum(element+1:last_element) < max_global then
	        break // mark as done
	      end if
	   end
	 end for
	*/

	input := []float64{9, 6, 7, 8, 10, 12, 11, 4, 20, 21, 22}

	sort.SliceStable(input, func(i, j int) bool {
		return input[j] < input[i]
	})

	maxSum := float64(0)
	maxSumArr := []float64{}
	contigous := []float64{}

	for i, v := range input {

		if len(contigous) == 0 || contigous[len(contigous)-1]-v == 1 {
			contigous = append(contigous, v)
			continue
		}

		contigousSum := Sum(contigous)
		if maxSum == 0 || contigousSum > float64(maxSum) {
			maxSum = contigousSum
			maxSumArr = slices.Clone(contigous)
			contigous = []float64{}

			if i < len(input) || Sum(input[i+1:]) < maxSum {
				log.Println("Early termination...")
				break
			}
		}

	}

	contigousSum := Sum(contigous)
	if contigousSum > float64(maxSum) {
		maxSum = contigousSum
		maxSumArr = slices.Clone(contigous)
	}
	log.Println(maxSumArr, maxSum)
}

func Sum(input []float64) (total float64) {
	for _, v := range input {
		total += v
	}

	return total
}

func Solution9() {

	/*
	 for each element do
	    if check in set element+1 not exits then
	       assumes it is the max in the array
	       cont = element
	       for cont is exist in set do
	          sum += cont
	          cont = cont-1
	        end for
	     end if
	 end for
	*/

	var min, max, sum int
	input := []int{9, 6, 7, 8, 10, 12, 11, 4, 20, 21, 22, 23}

	set := map[int]bool{}

	for _, v := range input {
		set[v] = true
	}

	for _, v := range input {
		if !set[v+1] {
			var cMin, cMax, cSum int
			cMax = v

			el := v
			for set[el] {
				cSum += el
				el = el - 1
			}
			cMin = el + 1

			if sum == 0 || cSum > sum {
				max = cMax
				min = cMin
				sum = cSum
			}

		}
	}

	log.Println(min, max, sum)

}
