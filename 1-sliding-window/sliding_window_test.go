package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubArrayAvg(t *testing.T) {
	assert.Equal(t, []float32{2.2, 2.8, 2.4, 3.6, 2.8}, SubArrayAvg([]float32{1, 3, 2, 6, -1, 4, 1, 8, 2}, 5))
}

func TestSubArrayMaxSum(t *testing.T) {
	// 9 -> [3, 4]
	assert.Equal(t, 9, SubArrayMaxSum([]int{2, 1, 5, 1, 3, 2}, 3))

	// 7 -> [5, 1, 3]
	assert.Equal(t, 7, SubArrayMaxSum([]int{2, 3, 4, 1, 5}, 2))
}

func TestSmallestSubArrayWithGreaterSum(t *testing.T) {
	// 7 -> [5, 2]
	assert.Equal(t, 2, SmallestSubArrayWithGreaterSum([]int{2, 1, 5, 2, 3, 2}, 7))

	// 7 -> [8]
	assert.Equal(t, 1, SmallestSubArrayWithGreaterSum([]int{2, 1, 5, 2, 8}, 7))

	// 8 -> [3, 4, 1] or [1, 1, 6]
	assert.Equal(t, 3, SmallestSubArrayWithGreaterSum([]int{3, 4, 1, 1, 6}, 8))
}

func TestLongerSubStrWithKDistChars(t *testing.T) {
	// "araaci" -> "araa"
	assert.Equal(t, 4, LongerSubStrWithKDistChars("araaci", 2))

	// "araaci" -> "aa"
	assert.Equal(t, 2, LongerSubStrWithKDistChars("araaci", 1))

	// "cbbebi" -> "cbbeb" & "bbebi"
	assert.Equal(t, 5, LongerSubStrWithKDistChars("cbbebi", 3))

	// "cbbebi" -> "cbbebi"
	assert.Equal(t, 6, LongerSubStrWithKDistChars("cbbebi", 10))
}

func TestFruitsIntoBaskets(t *testing.T) {
	// ['C', 'A', 'C']
	assert.Equal(t, 3, FruitsIntoBaskets([]string{"A", "B", "C", "A", "C"}))

	// ['B', 'C', 'B', 'B', 'C']
	assert.Equal(t, 5, FruitsIntoBaskets([]string{"A", "B", "C", "B", "B", "C"}))
}

func TestLongestSubStrWithDistChars(t *testing.T) {
	// "abc"
	assert.Equal(t, 3, LongestSubStrWithDistChars("aabccbb"))

	// "ab"
	assert.Equal(t, 2, LongestSubStrWithDistChars("abbbb"))

	// "abc" & "cde"
	assert.Equal(t, 3, LongestSubStrWithDistChars("abccde"))
}
