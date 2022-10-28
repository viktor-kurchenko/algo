package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseLinkedList(t *testing.T) {
	l := ReverseLinkedList(&Node{
		Val: 1,
		Next: &Node{
			Val: 2,
			Next: &Node{
				Val:  3,
				Next: nil,
			},
		},
	})
	for i := 3; i > 0; i-- {
		assert.Equal(t, i, l.Val)
		l = l.Next
	}
}
