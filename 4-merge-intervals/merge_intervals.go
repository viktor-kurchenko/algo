package main

func MergeIntervals(intervals [][]int) [][]int {
	m := make(map[int]int)
	for _, pair := range intervals {
		start, end := pair[0], pair[1]
		match := false
		for s, e := range m {
			if (start >= s && start <= e) || (end >= s && end <= e) || (start <= s && end >= e) {
				match = true
				if end > e {
					m[s] = end
				}
				if start < s {
					tmp := m[s]
					delete(m, s)
					m[start] = tmp
				}
				break
			}
		}
		if !match {
			m[start] = end
		}
	}
	result := make([][]int, 0)
	for s, e := range m {
		result = append(result, []int{s, e})
	}
	return result
}
