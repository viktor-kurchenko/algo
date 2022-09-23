package main

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
	result := 1
	for start, end := 0, 1; end < len(a); end++ {
		if a[start] != a[end] {
			result++
			start = end
		}
	}
	return result
}
