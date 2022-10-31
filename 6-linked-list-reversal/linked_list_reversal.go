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
	start := l
	for i := 1; i < p-1 && start.Next != nil; i++ {
		start = start.Next
	}

	var prev *Node
	head := start.Next
	for i := p; i <= q && head.Next != nil; i++ {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}

	start.Next, start.Next.Next = prev, head
	return l
}

type Node struct {
	Val  int
	Next *Node
}
