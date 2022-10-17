package main

import "sort"

func MergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		pStart, pEnd := result[len(result)-1][0], result[len(result)-1][1]
		start, end := intervals[i][0], intervals[i][1]
		if start < pEnd {
			result[len(result)-1][0], result[len(result)-1][1] = min(start, pStart), max(end, pEnd)
			continue
		}
		result = append(result, []int{start, end})
	}
	return result
}

func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	found := false
	for i := 0; i < len(intervals); i++ {
		if !found && newInterval[0] < intervals[i][1] {
			result = append(result, []int{min(newInterval[0], intervals[i][0]), max(newInterval[1], intervals[i][1])})
			found = true
			continue
		}
		if found && intervals[i][0] < result[len(result)-1][1] {
			result[len(result)-1][1] = max(result[len(result)-1][1], intervals[i][1])
			continue
		}
		result = append(result, []int{intervals[i][0], intervals[i][1]})
	}
	return result
}

func IntervalsIntersection(arr1, arr2 [][]int) [][]int {
	result := make([][]int, 0)
	for p1, p2 := 0, 0; p1 < len(arr1) && p2 < len(arr2); {
		i1, i2 := arr1[p1], arr2[p2]
		cmp := comp(i1, i2)
		if cmp == -1 {
			p2++
			continue
		}
		if cmp == 1 {
			p1++
			continue
		}
		result = append(result, []int{max(i1[0], i2[0]), min(i1[1], i2[1])})
		if i1[1] > i2[1] {
			p2++
			continue
		}
		if i1[1] < i2[1] {
			p1++
			continue
		}
		p1++
		p2++
	}
	return result
}

// helpers

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// -1: p1>p2, 0: intersection, 1: p1<p2
func comp(p1, p2 []int) int {
	if p1[0] > p2[1] {
		return -1
	}
	if p2[0] > p1[1] {
		return 1
	}
	return 0
}
