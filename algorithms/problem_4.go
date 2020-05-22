package algorithms

import (
	"fmt"
	"math"
)

func selectTop(tom []int, ti int, jerry []int, ji int) int {
	if ti == 0 && ji == 0 {
		return math.MinInt64
	}

	if ti == 0 {
		return jerry[ji-1]
	}

	if ji == 0 {
		return tom[ti-1]
	}
	tv, jv := tom[ti-1], jerry[ji-1]

	if tv < jv {
		return jv
	}

	return tv
}

func selectButtom(tom []int, ti int, jerry []int, ji int) int {
	if ti >= len(tom) && ji >= len(jerry) {
		return math.MaxInt64
	}

	if ti >= len(tom) {
		return jerry[ji]
	}

	if ji >= len(jerry) {
		return tom[ti]
	}

	tv, jv := tom[ti], jerry[ji]

	if tv < jv {
		return tv
	}

	return jv
}

// The key of this algorithm is to find the median index dividing points for both lists,
// and if we find the point for one list, we move forward to another list to do the same thing.
func findMedianSortedArrays(tom []int, jerry []int) float64 {
	if len(tom) == 0 {
		return findMedian(jerry)
	}

	if len(jerry) == 0 {
		return findMedian(tom)
	}

	ts, js := len(tom), len(jerry)

	totalSize := ts + js
	expectedIndex := totalSize / 2

	ss, se := 0, ts
	ls, le := 0, js
	candidate := tom[ts-1]

	var tomMedianIndex, jerryMedianIndex int

	for {
		sm := indexOf(tom, ss, se, candidate)
		lm := indexOf(jerry, ls, le, candidate)
		tomMedianIndex, jerryMedianIndex = sm, lm

		leftPartSize := sm + lm

		fmt.Printf("tom:%v,%d - %d - %d, jerry: %v, %d - %d - %d. candidate; %d, leftPart: %d Expected index: %d\n", tom, ss, sm, se, jerry, ls, lm, le, candidate, leftPartSize, expectedIndex)
		// We are just in the right place in both lists
		if leftPartSize == expectedIndex {

			break
		}

		var nextCandidate int

		// If there's still not enough elements on the left parts, move forward to the right. Otherwise move toward to the left
		if leftPartSize < expectedIndex {
			if se-sm > 1 {
				nextCandidateIndexStep := (se - sm) / 2
				nextCandidate = tom[sm+nextCandidateIndexStep]
				if nextCandidate == candidate {
					ss = sm + int(math.Max(float64(nextCandidateIndexStep/2), 1)) // If next candidate equals the previous one, it means all the elements in tom[sm:sm+(se-sm)] equal, we must make ss move forward, otherwise it will stay here forever
				} else {
					ss = sm
				}
				ls = lm
			} else {
				nextCandidateIndexStep := (le - lm) / 2
				nextCandidate = jerry[lm+nextCandidateIndexStep]

				if nextCandidate == candidate {
					ls = lm + int(math.Max(float64(nextCandidateIndexStep/2), 1))
				} else {
					ls = lm
				}
				ss = sm
			}
		} else {
			if se != sm {
				nextCandidate = tom[ss+(sm-ss)/2]
			} else {
				nextCandidate = jerry[ls+(lm-lm)/2]
			}
			se = sm
			le = lm
		}
		candidate = nextCandidate
	}

	if totalSize%2 == 0 {
		return float64(selectTop(tom, tomMedianIndex, jerry, jerryMedianIndex)+selectButtom(tom, tomMedianIndex, jerry, jerryMedianIndex)) / 2.0
	} else {
		return float64(selectButtom(tom, tomMedianIndex, jerry, jerryMedianIndex))
	}
}

// findMedian: Find the median number of a lsp
func findMedian(lst []int) float64 {
	size := len(lst)

	if size == 0 {
		return float64(math.MinInt64)
	}

	r := size / 2

	if size%2 == 0 {
		return float64(lst[r-1]+lst[r]) / 2.0
	}

	return float64(lst[r])
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
