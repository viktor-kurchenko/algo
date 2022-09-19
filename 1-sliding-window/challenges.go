package main

import "fmt"

func main() {
	subArrayMaxSum()
}

//Array: [1, 3, 2, 6, -1, 4, 1, 8, 2], K=5
//Output: [2.2, 2.8, 2.4, 3.6, 2.8]
func subArrayAvg() {
	array := []float32{1, 3, 2, 6, -1, 4, 1, 8, 2}
	const k = 5
	result := make([]float32, 0, len(array)-k+1)

	var sum float32
	for start, end := 0, 0; end < len(array); end++ {
		sum += array[end]
		if end < k-1 {
			continue
		}
		result = append(result, sum/k)
		sum -= array[start]
		start++
	}
	fmt.Printf("result: %v", result)
}

//Input: [2, 3, 4, 1, 5], k=2
//Output: 7 -> [3, 4]
func subArrayMaxSum() {
	array := []int{2, 3, 4, 1, 5}
	const k = 2
	var result, sum int
	for start, end := 0, 0; end < len(array); end++ {
		sum += array[end]
		if end < k-1 {
			continue
		}
		if result < sum {
			result = sum
		}
		sum -= array[start]
		start++
	}
	fmt.Printf("result: %d", result)
}
