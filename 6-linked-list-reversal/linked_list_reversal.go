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

type Node struct {
	Val  int
	Next *Node
}
