package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseLinkedList(t *testing.T) {
	original := &Node{
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
	expected := &Node{
		val: 5,
		next: &Node{
			val: 4,
			next: &Node{
				val: 3,
				next: &Node{
					val: 2,
					next: &Node{
						val: 1,
					},
				},
			},
		},
	}
	actual := ReverseLinkedList(original)
	for expected.next != nil && actual.next != nil {
		assert.Equal(t, expected.val, actual.val)
		expected, actual = expected.next, actual.next
	}
}

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
	assert.Equal(t, 3, MiddleLinkedList(&list).val)
}

func TestPalindromeLinkedList(t *testing.T) {
	list := Node{
		val: 1,
		next: &Node{
			val: 2,
			next: &Node{
				val: 4,
				next: &Node{
					val: 6,
					next: &Node{
						val: 4,
						next: &Node{
							val: 2,
							next: &Node{
								val: 1,
							},
						},
					},
				},
			},
		},
	}
	assert.True(t, PalindromeLinkedList(&list))

	list = Node{
		val: 2,
		next: &Node{
			val: 4,
			next: &Node{
				val: 6,
				next: &Node{
					val: 4,
					next: &Node{
						val: 2,
						next: &Node{
							val: 2,
						},
					},
				},
			},
		},
	}
	assert.False(t, PalindromeLinkedList(&list))
}

func TestRearrangeLinkedList(t *testing.T) {
	original := &Node{
		val: 2,
		next: &Node{
			val: 4,
			next: &Node{
				val: 6,
				next: &Node{
					val: 8,
					next: &Node{
						val: 10,
						next: &Node{
							val: 12,
						},
					},
				},
			},
		},
	}
	expected := &Node{
		val: 2,
		next: &Node{
			val: 12,
			next: &Node{
				val: 4,
				next: &Node{
					val: 10,
					next: &Node{
						val: 6,
						next: &Node{
							val: 8,
						},
					},
				},
			},
		},
	}
	actual := RearrangeLinkedList(original)
	assert.True(t, linkedListsEqual(expected, actual))

	original = &Node{
		val: 2,
		next: &Node{
			val: 4,
			next: &Node{
				val: 6,
				next: &Node{
					val: 8,
					next: &Node{
						val: 10,
					},
				},
			},
		},
	}
	expected = &Node{
		val: 2,
		next: &Node{
			val: 10,
			next: &Node{
				val: 4,
				next: &Node{
					val: 8,
					next: &Node{
						val: 6,
					},
				},
			},
		},
	}
	actual = RearrangeLinkedList(original)
	assert.True(t, linkedListsEqual(expected, actual))
}

// helpers

func linkedListsEqual(list1, list2 *Node) bool {
	for list1.next != nil {
		if list1.val != list2.val {
			return false
		}
		list1, list2 = list1.next, list2.next
	}
	return true
}
