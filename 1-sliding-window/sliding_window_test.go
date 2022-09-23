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

func TestLongestSubStrWithDistCharsAfterReplacement(t *testing.T) {
	// "bbbbb"
	assert.Equal(t, 5, LongestSubStrWithDistCharsAfterReplacement("aabccbb", 2))

	// "bbbb"
	assert.Equal(t, 4, LongestSubStrWithDistCharsAfterReplacement("abbcb", 1))

	// "ccc"
	assert.Equal(t, 3, LongestSubStrWithDistCharsAfterReplacement("abccde", 1))
}

func TestLongestSubArrayWithOnesAfterReplacement(t *testing.T) {
	// Replace the '0' at index 5 and 8 to have the longest contiguous subarray of 1s having length 6.
	assert.Equal(t, 6, LongestSubArrayWithOnesAfterReplacement([]int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}, 2))

	// Replace the '0' at index 6, 9, and 10 to have the longest contiguous subarray of 1s having length 9.
	assert.Equal(t, 9, LongestSubArrayWithOnesAfterReplacement([]int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}, 3))
}

func TestPermutationInString(t *testing.T) {
	// The string contains "bca" which is a permutation of the given pattern.
	assert.True(t, PermutationInString("oidbcaf", "abc"))

	assert.False(t, PermutationInString("odicf", "dc"))

	// Both the string and the pattern are a permutation of each other.
	assert.True(t, PermutationInString("bcdxabcdy", "bcdyabcdx"))

	// The string contains "acb" which is a permutation of the given pattern.
	assert.True(t, PermutationInString("aaacb", "abc"))
}

func TestStrAnagrams(t *testing.T) {
	// "pq" and "qp"
	assert.Equal(t, []int{1, 2}, StrAnagrams("ppqp", "pq"))

	// "bca", "cab", and "abc"
	assert.Equal(t, []int{2, 3, 4}, StrAnagrams("abbcabc", "abc"))
}

func TestSmallestWindowWithSubStr(t *testing.T) {
	assert.Equal(t, "abdec", SmallestWindowWithSubStr("aabdec", "abc"))

	assert.Equal(t, "aabdec", SmallestWindowWithSubStr("aabdec", "abac"))

	assert.Equal(t, "bca", SmallestWindowWithSubStr("abdbca", "abc"))

	assert.Equal(t, "", SmallestWindowWithSubStr("adcad", "abc"))
}

func TestWordsConcatenation(t *testing.T) {
	// The two substring containing both the words are "catfox" & "foxcat".
	assert.Equal(t, []int{0, 3}, WordsConcatenation("catfoxcat", []string{"cat", "fox"}))

	// The only substring containing both the words is "catfox".
	assert.Equal(t, []int{3}, WordsConcatenation("catcatfoxfox", []string{"cat", "fox"}))
}
