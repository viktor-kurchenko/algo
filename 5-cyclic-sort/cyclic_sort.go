package main

import "sort"

func CyclicSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for a[i] != i+1 {
			a[i], a[a[i]-1] = a[a[i]-1], a[i]
		}
	}
	return a
}

func MissingNumber(a []int) int {
	for i := 0; i < len(a); i++ {
		for a[i] != i {
			if a[i] == len(a) {
				break
			}
			a[i], a[a[i]] = a[a[i]], a[i]
		}
	}
	for i := 0; i < len(a); i++ {
		if a[i] != i {
			return i
		}
	}
	return -1
}

func AllMissingNumbers(a []int) []int {
	for i := 0; i < len(a); i++ {
		for a[i] != i+1 && a[i] != a[a[i]-1] {
			a[i], a[a[i]-1] = a[a[i]-1], a[i]
		}
	}
	result := make([]int, 0)
	for i := 0; i < len(a); i++ {
		if a[i] != i+1 {
			result = append(result, i+1)
		}
	}
	return result
}

func FindDuplicate(a []int) int {
	for i := 0; i < len(a); i++ {
		for a[i] != i+1 {
			if a[i] == a[a[i]-1] {
				return a[i]
			}
			a[i], a[a[i]-1] = a[a[i]-1], a[i]
		}
	}
	return -1
}

func FindAllDuplicates(a []int) []int {
	result := make([]int, 0)
	for i := 0; i < len(a); i++ {
		for a[i] != i+1 {
			if a[i] == a[a[i]-1] {
				result = append(result, a[i])
				break
			}
			a[i], a[a[i]-1] = a[a[i]-1], a[i]
		}
	}
	sort.Ints(result)
	return result
}

func FindCorruptPair(a []int) []int {
	for i := 0; i < len(a); i++ {
		for a[i] != i+1 && a[i] != a[a[i]-1] {
			a[i], a[a[i]-1] = a[a[i]-1], a[i]
		}
	}
	for i := 0; i < len(a); i++ {
		if a[i] != i+1 {
			return []int{a[i], i + 1}
		}
	}
	return nil
}

func FindSmallestMissingPositiveNumber(a []int) int {
	for i := 0; i < len(a); i++ {
		for a[i] != i+1 {
			if a[i] < 1 || a[i] > len(a) || a[i] == a[a[i]-1] {
				break
			}
			a[i], a[a[i]-1] = a[a[i]-1], a[i]
		}
	}
	for i := 0; i < len(a); i++ {
		if a[i] != i+1 {
			return i + 1
		}
	}
	return -1
}

func FindFirstKSmallestMissingPositiveNumbers(a []int, k int) []int {
	large := make(map[int]struct{})
	for i := 0; i < len(a); i++ {
		if a[i] < 1 {
			continue
		}
		for a[i] != i+1 || a[i] > 0 {
			if a[i] > len(a) {
				large[a[i]], a[i] = struct{}{}, 0
				break
			}
			if a[i] == a[a[i]-1] {
				break
			}
			a[i], a[a[i]-1] = a[a[i]-1], a[i]
		}
	}
	result := make([]int, 0)
	for i := 0; i < len(a) && len(result) < k; i++ {
		if a[i] == i+1 {
			continue
		}
		result = append(result, i+1)
	}
	for i := len(a) + 1; len(result) < k; i++ {
		if _, ok := large[i]; ok {
			continue
		}
		result = append(result, i)
	}
	return result
}
