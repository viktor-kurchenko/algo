package main

import "sort"

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
		for start, end := i+1, len(a)-1; start < end; {
			y := a[start]
			z := a[end]
			s := y + z + x
			if s == 0 {
				if len(result) > 0 {
					p := result[len(result)-1]
					if p[0] == x && p[1] == y && p[2] == z {
						start++
						continue
					}
				}
				result = append(result, []int{x, y, z})
			}
			if s <= 0 {
				start++
				continue
			}
			end--
		}
	}
	return result
}
