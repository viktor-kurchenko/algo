package main

func CyclicSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for a[i] != i+1 {
			a[i], a[a[i]-1] = a[a[i]-1], a[i]
		}
	}
	return a
}
