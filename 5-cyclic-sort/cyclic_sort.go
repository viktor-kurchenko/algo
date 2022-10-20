package main

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
