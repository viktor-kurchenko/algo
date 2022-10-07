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

func TestStartLinkedListCycle(t *testing.T) {
	node6 := &Node{val: 6}
	node3 := &Node{
		val: 3,
		next: &Node{
			val: 4,
			next: &Node{
				val:  5,
				next: node6,
			},
		},
	}
	list := Node{
		val: 1,
		next: &Node{
			val: 2,
			next: &Node{
				val:  3,
				next: node3,
			},
		},
	}
	node6.next = list.next.next
	assert.Equal(t, node3.val, StartLinkedListCycle(&list).val)
}

func TestHappyNumber(t *testing.T) {
	assert.True(t, HappyNumber(23))
	assert.False(t, HappyNumber(12))
}

func TestMiddleLinkedList(t *testing.T) {
	list := Node{
		val: 1,
		next: &Node{
			val: 2,
			next: &Node{
				val: 3,
				next: &Node{
					val: 4,
					next: &Node{
						val: 5,
					},
				},
			},
		},
	}
	assert.Equal(t, 3, MiddleLinkedList(&list))
}
