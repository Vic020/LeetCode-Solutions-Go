package main

import (
	"testing"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestTwoSum(t *testing.T) {
	result := twoSum([]int{1, 2, 3}, 4)
	target := []int{0, 2}
	if !Equal(result, target) {
		t.Errorf("Two Sum incorrect! got %v, want %v. ", result, target)
	}
}
