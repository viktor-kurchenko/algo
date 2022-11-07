package main

func ReverseLinkedList(l *Node) *Node {
	if l.Next == nil {
		return l
	}
	var prev *Node
	node := l
	for node.Next != nil {
		next := node.Next
		node.Next = prev
		prev, node = node, next
	}
	node.Next = prev
	return node
}

func ReverseSubList(l *Node, p, q int) *Node {
	var prev *Node
	node := l
	for i := 1; i < p; i++ {
		prev = node
		node = node.Next
	}

	start := prev
	subListEnd := node

	for i := p; i <= q; i++ {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}

	if start == nil {
		start = prev
		subListEnd.Next = node
	} else {
		start.Next.Next = node
		start.Next = prev
	}
	return start
}

type Node struct {
	Val  int
	Next *Node
}
