package main

import "fmt"

func main() {
	//Array: [1, 3, 2, 6, -1, 4, 1, 8, 2], K=5
	//Output: [2.2, 2.8, 2.4, 3.6, 2.8]
	array := []float32{1, 3, 2, 6, -1, 4, 1, 8, 2}
	const k = 5
	result := make([]float32, 0, len(array)-k+1)

	var sum float32
	for start, end := 0, 0; end < len(array); end++ {
		if end-start < k {
			sum += array[end]
			continue
		}
		result = append(result, sum/k)
		sum -= array[start]
		start++
		sum += array[end]
	}
	result = append(result, sum/k)
	fmt.Printf("result: %v", result)
}
