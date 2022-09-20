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
