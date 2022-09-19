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
	var start, sum int
	for end := 0; end < len(a); end++ {
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
