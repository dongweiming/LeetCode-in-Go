package problem0056

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	merged := &[]Interval{}
	for _, i := range intervals {
		size := len(*merged)
		if len(*merged) == 0 || (*merged)[size-1].End < i.Start {
			*merged = append(*merged, i)
		} else if (*merged)[size-1].End < i.End {
			(*merged)[size-1].End = i.End
		}
	}
	return *merged
}

// For test
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
