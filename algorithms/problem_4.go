package algorithms

import "math"

func findMedianSortedArrays(tom []int, jerry []int) float64 {
	medianEleIndices := medianIndices(tom, jerry)

	if len(medianEleIndices) == 0 {
		// Use math.MaxFloat64 as error mark
		return math.MaxFloat64
	}
	firstMedianElementIndex := medianEleIndices[0]
	var candidate int

	if len(tom) > firstMedianElementIndex {
		candidate = tom[firstMedianElementIndex]
	} else {
		candidate = jerry[firstMedianElementIndex]
	}

	startTom, medianTom, endTom := 0, 0, len(tom)
	startJerry, medianJerry, endJerry := 0, 0, len(jerry)
	for {
	}

	return 0
}

// Find the indices of the median elements
// Let s = sum(len(tom) + len(jerry)), the indices of median elements are:
// - if s is even: [s/2 - 1, s/2]
// - if s is odd: [s/2]
func medianIndices(tom []int, jerry []int) []int {
	s := len(tom) + len(jerry)
	if s == 0 {
		return []int{}
	}
	index := s / 2
	if s%2 == 0 {
		return []int{index - 1, index}
	} else {
		return []int{index}
	}
}

// Given a orderred list, find the index of the target number between lst[start_inclusive, end_exclusive)
// Index of the target number means: if we insert this number into the given list, which index will it be put
// If the target number is already in the lst, return the left-most index.
func indexOf(lst []int, startInclusive, endExclusive int, target int) int {
	if startInclusive == endExclusive {
		return startInclusive
	}

	if target < lst[startInclusive] {
		return startInclusive
	}

	if target > lst[endExclusive-1] {
		return endExclusive
	}

	middleIndex := startInclusive + (endExclusive-startInclusive)/2

	ele := lst[middleIndex]

	if ele < target {
		return indexOf(lst, middleIndex+1, endExclusive, target)
	}

	return indexOf(lst, startInclusive, middleIndex, target)
}
