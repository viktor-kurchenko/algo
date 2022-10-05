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

func LinkedListCycleLength(list *Node) int {
	slow, fast := list, list
	for {
		slow, fast = slow.next, fast.next.next
		if slow == fast {
			break
		}
	}
	n := slow.next
	i := 1
	for n != slow {
		i++
		n = n.next
	}
	return i
}

// helpers

type Node struct {
	val  int
	next *Node
}
