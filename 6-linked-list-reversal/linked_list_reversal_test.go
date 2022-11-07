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

func TestReverseSubList(t *testing.T) {
	actual := ReverseSubList(&Node{
		Val: 1,
		Next: &Node{
			Val: 2,
			Next: &Node{
				Val: 3,
				Next: &Node{
					Val: 4,
					Next: &Node{
						Val: 5,
					},
				},
			},
		},
	}, 1, 4)
	expected := &Node{
		Val: 4,
		Next: &Node{
			Val: 3,
			Next: &Node{
				Val: 2,
				Next: &Node{
					Val: 1,
					Next: &Node{
						Val: 5,
					},
				},
			},
		},
	}
	for i := 0; i < 5; i++ {
		assert.Equal(t, expected.Val, actual.Val)
		actual, expected = actual.Next, expected.Next
	}
}
