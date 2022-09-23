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
