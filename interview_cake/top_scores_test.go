package main

import (
	"reflect"
	"testing"
)

func TestSortScores(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      int
		expected []int
	}{
		{[]int{}, 0, []int{}},
		{[]int{37, 89, 41, 65, 91, 53}, 100, []int{91, 89, 65, 53, 41, 37}},
		{[]int{37, 89, 41, 65, 91, 53, 65}, 100, []int{91, 89, 65, 65, 53, 41, 37}},
	}

	for _, tt := range tests {
		result := sortScores(tt.in1, tt.in2)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

// sortScores takes an unsorted list of integers and the highest possible score
// and returns a sorted list. Since we build a list of socres where indices
// represent scores and values represent how many time the score appears, the
// space complexity is O(n). The time complexity is O(n).
func sortScores(scores []int, highestPossibleScore int) []int {
	counts := make([]int, highestPossibleScore+1)
	out := make([]int, 0)

	// populate the scores where indeices represent scores and values represent
	// how many time the score appears.
	for _, score := range scores {
		counts[score] += 1
	}

	// iterate backward and add the score to the output list so that it is
	// sorted from highest to lowest.
	for i := len(counts) - 1; i > -1; i-- {
		count := counts[i]

		for j := 0; j < count; j++ {
			out = append(out, i)
		}
	}

	return out
}
