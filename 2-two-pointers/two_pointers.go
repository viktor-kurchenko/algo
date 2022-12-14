package main

import (
	"math"
	"sort"
)

func PairWithTargetSum(a []int, s int) []int {
	for start, end := 0, len(a)-1; start < end; {
		sum := a[start] + a[end]
		if sum == s {
			return []int{start, end}
		}
		if sum > s {
			end--
			continue
		}
		start++
	}
	return nil
}

func RemoveDuplicates(a []int) int {
	next := 1
	for i := 0; i < len(a); i++ {
		if a[next-1] != a[i] {
			a[next] = a[i]
			next++
		}
	}
	return next
}

func RemoveKey(a []int, k int) int {
	next := 0
	for i := 0; i < len(a); i++ {
		if a[i] != k {
			a[next] = a[i]
			next++
		}
	}
	return next
}

func SquaringSortedArray(a []int) []int {
	result := make([]int, len(a))
	start, end := 0, len(a)-1
	index := len(a) - 1
	for start < end {
		s1 := a[start] * a[start]
		s2 := a[end] * a[end]
		if s1 <= s2 {
			result[index] = s2
			end--
		} else {
			result[index] = s1
			start++
		}
		index--
	}
	return result
}

func TripletSum2Zero(a []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(a)
	for i := 0; i < len(a)-2; i++ {
		x := a[i]
		if i > 0 && x == a[i-1] {
			continue
		}
		for start, end := i+1, len(a)-1; start < end; {
			y := a[start]
			z := a[end]
			if -x == y+z {
				result = append(result, []int{x, y, z})
				start++
				end--
				for start < end && a[start] == a[start-1] {
					start++
				}
				for start < end && a[end] == a[end+1] {
					end--
				}
				continue
			}
			if -x > y+z {
				start++
			} else {
				end--
			}
		}
	}
	return result
}

func TripletSumCloseToTarget(a []int, t int) int {
	sort.Ints(a)
	result := math.MaxInt
	for i := 0; i < len(a)-2; i++ {
		for start, end := i+1, len(a)-1; start < end; {
			s := a[i] + a[start] + a[end]
			if s == t {
				return s
			}
			if abs(t-s) < abs(t-result) {
				result = s
			}
			if s < t {
				start++
				continue
			}
			end--
		}
	}
	return result
}

func TripletWithSmallerSum(a []int, t int) int {
	sort.Ints(a)
	result := 0
	for i := 0; i < len(a)-2; i++ {
		for start, end := i+1, len(a)-1; start < end; {
			s := a[i] + a[start] + a[end]
			if s < t {
				result += end - start
				start++
				continue
			}
			end--
		}
	}
	return result
}

func SubArraysWithProductLessThanTarget(a []int, t int) [][]int {
	result := make([][]int, 0)
	p := 1
	for start, end := 0, 0; end < len(a); end++ {
		if a[end] < t {
			result = append(result, []int{a[end]})
		}
		p *= a[end]
		if start == end {
			continue
		}
		if p < t {
			r := make([]int, 0)
			for j := start; j <= end; j++ {
				r = append(r, a[j])
			}
			result = append(result, r)
			continue
		}
		for start < end-1 {
			p /= a[start]
			start++
			if p < t {
				r := make([]int, 0)
				for j := start; j <= end; j++ {
					r = append(r, a[j])
				}
				result = append(result, r)
			}
		}
	}
	return result
}

func DutchNationFlagProblem(a []int) []int {
	for low, i, high := 0, 0, len(a)-1; i <= high; {
		switch a[i] {
		case 0:
			a[i], a[low] = a[low], 0
			i++
			low++
		case 1:
			i++
		case 2:
			a[i], a[high] = a[high], 2
			high--
		}
	}
	return a
}

func QuadrupleSum2Target(a []int, t int) [][]int {
	result := make([][]int, 0)
	sort.Ints(a)
	for i := 0; i < len(a)-3; i++ {
		if i > 0 && a[i] == a[i-1] {
			continue
		}
		for j := i + 1; j < len(a)-2; j++ {
			if j > i+1 && a[j] == a[j-1] {
				continue
			}
			for s, e := j+1, len(a)-1; s < e; {
				sum := a[i] + a[j] + a[s] + a[e]
				if sum > t {
					e--
					for s < e && a[e] == a[e+1] {
						e--
					}
					continue
				}
				if sum == t {
					result = append(result, []int{a[i], a[j], a[s], a[e]})
				}
				s++
				for s < e && a[s] == a[s-1] {
					s++
				}
			}
		}
	}
	return result
}

func ComparingStringsWithBackspaces(s1, s2 string) bool {
	for i1, i2 := len(s1)-1, len(s2)-1; i1 >= 0 && i2 >= 0; i1, i2 = i1-1, i2-1 {
		r := 0
		for s1[i1] == '#' || r > 0 {
			if s1[i1] == '#' {
				r++
				i1--
				continue
			}
			i1--
			r--
		}
		for s2[i2] == '#' || r > 0 {
			if s2[i2] == '#' {
				r++
				i2--
				continue
			}
			i2--
			r--
		}
		if s1[i1] != s2[i2] {
			return false
		}
		if i1 == 0 && i2 == 0 {
			return true
		}
	}
	return false
}

func MinWindow2Sort(a []int) int {
	minIdx, maxIdx := 0, 0
	for max, i := a[0], 1; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
		if a[i] < a[i-1] || a[i] < max {
			minIdx = i
		}
	}
	for min, i := len(a)-1, len(a)-2; i >= 0; i-- {
		if min > a[i] {
			min = a[i]
		}
		if a[i] > a[i+1] || a[i] > min {
			maxIdx = i
		}
	}
	if maxIdx == 0 && minIdx == 0 {
		return 0
	}
	return abs(maxIdx-minIdx) + 1
}

// helpers

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
