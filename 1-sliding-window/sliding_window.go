package main

import "reflect"

func SubArrayAvg(a []float32, k int) []float32 {
	result := make([]float32, 0, len(a)-k+1)
	var sum float32
	for start, end := 0, 0; end < len(a); end++ {
		sum += a[end]
		if end < k-1 {
			continue
		}
		result = append(result, sum/float32(k))
		sum -= a[start]
		start++
	}
	return result
}

func SubArrayMaxSum(a []int, k int) int {
	var result, sum int
	for start, end := 0, 0; end < len(a); end++ {
		sum += a[end]
		if end < k-1 {
			continue
		}
		if result < sum {
			result = sum
		}
		sum -= a[start]
		start++
	}
	return result
}

func SmallestSubArrayWithGreaterSum(a []int, s int) int {
	result := len(a)
	var sum int
	for start, end := 0, 0; end < len(a); end++ {
		sum += a[end]
		for sum >= s {
			size := end - start + 1
			if result > size {
				result = size
			}
			sum -= a[start]
			start++
		}
	}
	return result
}

func LongerSubStrWithKDistChars(s string, k int) int {
	var result int
	m := make(map[uint8]int)
	for start, end := 0, 0; end < len(s); end++ {
		m[s[end]] = m[s[end]] + 1
		for ; len(m) > k; start++ {
			count := m[s[start]]
			if count == 1 {
				delete(m, s[start])
				continue
			}
			m[s[start]] = count - 1
		}
		size := end - start + 1
		if result < size {
			result = size
		}
	}
	return result
}

func FruitsIntoBaskets(a []string) int {
	const k = 2
	var result int
	m := make(map[string]int)
	for start, end := 0, 0; end < len(a); end++ {
		m[a[end]] = m[a[end]] + 1
		for ; len(m) > k; start++ {
			c := a[start]
			count := m[c]
			if count == 1 {
				delete(m, c)
				continue
			}
			m[c] = count - 1
		}
		size := end - start + 1
		if result < size {
			result = size
		}
	}
	return result
}

func LongestSubStrWithDistChars(s string) int {
	var result int
	m := make(map[uint8]int)
	for start, end := 0, 0; end < len(s); end++ {
		c := s[end]
		if p, ok := m[c]; ok {
			start = p + 1
		}
		m[c] = end
		size := end - start + 1
		if result < size {
			result = size
		}
	}
	return result
}

func LongestSubStrWithDistCharsAfterReplacement(s string, k int) int {
	var result, maxRep int
	m := make(map[uint8]int)
	for start, end := 0, 0; end < len(s); end++ {
		c := s[end]
		m[c] = m[c] + 1
		if maxRep < m[c] {
			maxRep = m[c]
		}
		if end-start+1-maxRep > k {
			m[s[start]] = m[s[start]] - 1
			start++
		}
		size := end - start + 1
		if result < size {
			result = size
		}
	}
	return result
}

func LongestSubArrayWithOnesAfterReplacement(a []int, k int) int {
	var result, maxRep int
	for start, end := 0, 0; end < len(a); end++ {
		if a[end] == 1 {
			maxRep++
		}
		if end-start+1-maxRep > k {
			if a[start] == 1 {
				maxRep--
			}
			start++
		}
		size := end - start + 1
		if result < size {
			result = size
		}
	}
	return result
}

func PermutationInString(s, p string) bool {
	pm := make(map[uint8]int)
	for i := range p {
		c := p[i]
		pm[c] = pm[c] + 1
	}
	m := make(map[uint8]int)
	for start, end := 0, 0; end < len(s); end++ {
		m[s[end]] = m[s[end]] + 1
		if end-start+1 < len(p) {
			continue
		}
		if reflect.DeepEqual(m, pm) {
			return true
		}
		c := s[start]
		if m[c] == 1 {
			delete(m, c)
		} else {
			m[c] = m[c] - 1
		}
		start++
	}
	return false
}
