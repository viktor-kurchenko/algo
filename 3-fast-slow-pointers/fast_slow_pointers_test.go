package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListCycle(t *testing.T) {
	node6 := &Node{val: 6}
	list := Node{
		val: 1,
		next: &Node{
			val: 2,
			next: &Node{
				val: 3,
				next: &Node{
					val: 4,
					next: &Node{
						val:  5,
						next: node6,
					},
				},
			},
		},
	}
	assert.False(t, LinkedListCycle(&list))

	node6.next = list.next.next
	assert.True(t, LinkedListCycle(&list))
}

func TestLinkedListCycleLength(t *testing.T) {
	node6 := &Node{val: 6}
	list := Node{
		val: 1,
		next: &Node{
			val: 2,
			next: &Node{
				val: 3,
				next: &Node{
					val: 4,
					next: &Node{
						val:  5,
						next: node6,
					},
				},
			},
		},
	}
	node6.next = list.next.next
	assert.Equal(t, 4, LinkedListCycleLength(&list))
}
