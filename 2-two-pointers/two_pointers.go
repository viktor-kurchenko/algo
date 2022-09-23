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
