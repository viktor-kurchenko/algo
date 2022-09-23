package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairWithTargetSum(t *testing.T) {
	// The numbers at index 1 and 3 add up to 6: 2+4=6
	assert.Equal(t, []int{1, 3}, PairWithTargetSum([]int{1, 2, 3, 4, 6}, 6))

	// The numbers at index 0 and 2 add up to 11: 2+9=11
	assert.Equal(t, []int{0, 2}, PairWithTargetSum([]int{2, 5, 9, 11}, 11))
}
