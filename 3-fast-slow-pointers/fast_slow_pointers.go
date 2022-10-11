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

func HappyNumber(n int) bool {
	m := make(map[int]struct{})
	m[n] = struct{}{}
	for {
		n = num2SquareSum(n)
		if n == 1 {
			return true
		}
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = struct{}{}
	}
}

func MiddleLinkedList(list *Node) *Node {
	slow, fast := list, list.next
	for fast != nil && fast.next != nil {
		slow, fast = slow.next, fast.next.next
	}
	return slow
}

func PalindromeLinkedList(list *Node) bool {
	mid, fast := list, list
	for fast.next != nil && fast.next.next != nil {
		mid, fast = mid.next, fast.next.next
	}
	// reverse second half
	mid = ReverseLinkedList(mid)
	first, second := list, mid
	equal := true
	for first.next != nil && second.next != nil {
		if first.val != second.val {
			equal = false
			break
		}
		first, second = first.next, second.next
	}
	// reverse second half again
	ReverseLinkedList(mid)
	return equal
}

func RearrangeLinkedList(list *Node) *Node {
	mid, fast := list, list
	for fast.next != nil && fast.next.next != nil {
		mid, fast = mid.next, fast.next.next
	}
	start, mid := list, ReverseLinkedList(mid)
	for start.next != nil && mid.next != nil {
		start, start.next = start.next, mid
		mid, mid.next = mid.next, start
	}
	return list
}

// helpers

type Node struct {
	val  int
	next *Node
}

func ReverseLinkedList(head *Node) *Node {
	var prev *Node
	for head != nil {
		next := head.next
		head.next = prev
		prev = head
		head = next
	}
	return prev
}

func num2SquareSum(n int) int {
	sum := 0
	for n > 0 {
		x := n % 10
		sum += x * x
		n /= 10
	}
	return sum
}
