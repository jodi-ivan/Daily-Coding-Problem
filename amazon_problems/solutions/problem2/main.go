package main

import (
	"strings"
)

func MaxSubstring(str string, k int, left, right int, freq map[byte]int, best string) string {
	// Debugging logs

	// Base case: Stop when `right` reaches end
	if right == len(str) {
		if len(freq) <= k && len(str[left:right]) > len(best) {
			best = str[left:right]
		}
		return best
	}

	// Expand the right side
	freq[str[right]]++

	// If too many distinct characters, shrink from left
	for len(freq) > k {
		freq[str[left]]--
		if freq[str[left]] == 0 {
			delete(freq, str[left]) // Remove from map when count is zero
		}
		left++
	}

	// Check if current window is the longest so far
	if len(str[left:right+1]) > len(best) {
		best = str[left : right+1]
	}

	// Recur with `right` expanded
	return MaxSubstring(str, k, left, right+1, freq, best)
}

func MaxSubstringLoop(str string, n int) string {
	var left, right int
	distinct := map[string]bool{}
	right = left + 1

	distinct[string(str[left])] = true
	distinct[string(str[right])] = true

	for right < len(str) {

		if len(distinct) > n {
			toRemove := string(str[left])
			left++

			if !strings.Contains(str[left:right-1], toRemove) {
				delete(distinct, toRemove)
				continue
			}
		}

		if ok := distinct[string(str[right])]; ok {
			right++ // expand
			continue
		} else {
			distinct[string(str[right])] = true
		}

		right++
	}

	return str[left:right]
}

func longestSubstringWithKDistinct(s string, k int) string {
	if k == 0 || len(s) == 0 {
		return ""
	}

	left, right := 0, 0
	freq := make(map[byte]int)
	maxLen := 0
	start := 0 // To track the starting index of the longest substring

	for right < len(s) {
		freq[s[right]]++

		// If we exceed k distinct characters, move left pointer
		for len(freq) > k {
			freq[s[left]]--
			if freq[s[left]] == 0 {
				delete(freq, s[left]) // Remove from map if count is 0
			}
			left++
		}

		// Update max length and starting index
		if right-left+1 > maxLen {
			maxLen = right - left + 1
			start = left
		}

		right++
	}

	return s[start : start+maxLen]
}

func main() {

}
