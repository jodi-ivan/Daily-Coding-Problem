package main

import (
	"encoding/json"
	"log"
	"strings"
)

/*

This problem was asked by Amazon.
There exists a staircase with N steps, and you can climb up either 1 or 2 steps
at a time. Given N, write a function that returns the number of unique ways you can
 climb the staircase. The order of the steps matters.
For example, if N is 4, then there are 5 unique ways:
•	1, 1, 1, 1
•	2, 1, 1
•	1, 2, 1
•	1, 1, 2
•	2, 2
What if, instead of being able to climb 1 or 2 steps at a time, you could
climb any number from a set of positive integers X? For example,
 if X = {1, 3, 5}, you could climb 1, 3, or 5 steps at a time*/

var STEPS = []int{1, 2, 3, 5}

func UnqieSteps(stairsteps int, index int, currWay []int, ways []string) []string {
	if index >= stairsteps {
		s, _ := json.Marshal(currWay)

		if index == stairsteps {
			return append(ways, strings.Trim(string(s), "[]"))
		}

		return ways
	}

	result := []string{}
	for _, v := range STEPS {
		steps := UnqieSteps(stairsteps, index+v, append(currWay, v), ways)
		result = append(result, steps...)
	}

	return result
}

func main() {
	res := UnqieSteps(6, 0, []int{}, []string{})
	raw, _ := json.MarshalIndent(res, "", "    ")
	log.Println(string(raw))
}
