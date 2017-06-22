package algo

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func tailPointer(current **Node) **Node {
	if (*current) == nil {
		return current
	}
	return tailPointer(&((*current).next))
}

func newNodeOf(val int) *Node {
	return &Node{value: val}
}

func (l *LinkedList) Add(value int) {
	tail := tailPointer(&l.head)

	*tail = newNodeOf(value)
}

func reverseRec(previous, current *Node) *Node {
	if current == nil {
		return previous
	}
	next := current.next
	current.next = previous
	return reverseRec(current, next)
}

func (l *LinkedList) Reverse() {
	l.head = reverseRec(nil, l.head)
}
