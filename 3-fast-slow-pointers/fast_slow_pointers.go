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

func StartLinkedListCycle(list *Node) *Node {
	p1, p2 := list.next, list.next.next
	for p1 != p2 {
		p1, p2 = p1.next, p2.next.next
	}
	l := 1
	p2 = p2.next
	for p1 != p2 {
		p2 = p2.next
		l++
	}
	p1, p2 = list, list
	for i := 0; i < l; i++ {
		p2 = p2.next
	}
	for p1 != p2 {
		p1, p2 = p1.next, p2.next
	}
	return p1
}

// helpers

type Node struct {
	val  int
	next *Node
}
