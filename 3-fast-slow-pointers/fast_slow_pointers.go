package main

func LinkedListCycle(list *Node) bool {
	slow, fast := list, list
	for slow.next != nil && fast.next.next != nil {
		slow, fast = slow.next, fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

// helpers

type Node struct {
	val  int
	next *Node
}
