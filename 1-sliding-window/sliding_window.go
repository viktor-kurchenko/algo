package main

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
		pm[p[i]] = pm[p[i]] + 1
	}
	matched := 0
	for start, end := 0, 0; end < len(s); end++ {
		if count, ok := pm[s[end]]; ok {
			pm[s[end]] = count - 1
			if count == 1 {
				matched++
			}
		}
		if end < len(p)-1 {
			continue
		}
		if matched == len(pm) {
			return true
		}
		if count, ok := pm[s[start]]; ok {
			pm[s[start]] = count + 1
			if count == 0 {
				matched--
			}
		}
		start++
	}
	return false
}

func StrAnagrams(s, a string) []int {
	m := make(map[uint8]int)
	for i := range a {
		m[a[i]] = m[a[i]] + 1
	}
	match := 0
	result := make([]int, 0)
	for start, end := 0, 0; end < len(s); end++ {
		if count, ok := m[s[end]]; ok {
			m[s[end]] = count - 1
			if count == 1 {
				match++
			}
		}
		if end < len(a)-1 {
			continue
		}
		if match == len(m) {
			result = append(result, start)
		}
		if count, ok := m[s[start]]; ok {
			m[s[start]] = count + 1
			if count == 0 {
				match--
			}
		}
		start++
	}
	return result
}

func SmallestWindowWithSubStr(s, p string) string {
	m := make(map[uint8]int)
	for i := range p {
		m[p[i]] = m[p[i]] + 1
	}
	result := ""
	matched := 0
	for start, end := 0, 0; end < len(s); end++ {
		if count, ok := m[s[end]]; ok {
			m[s[end]] = count - 1
			if count == 1 {
				matched++
			}
		}
		if end < len(p)-1 {
			continue
		}
		for matched == len(m) {
			if result == "" || len(result) > end-start+1 {
				result = s[start : end+1]
			}
			if count, ok := m[s[start]]; ok {
				m[s[start]] = count + 1
				if count == 0 {
					matched--
				}
			}
			start++
		}
	}
	return result
}
