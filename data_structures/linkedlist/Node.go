package linkedlist

type Node struct {
	Val      int
	Next     *Node
	Previous *Node
}

func NewNode(new_val int) *Node {
	return &Node{
		new_val,
		nil,
		nil,
	}
}

func (n *Node) updateVal(Val int) {
	n.Val = Val
}

func (n *Node) updateNext(Next *Node) {
	n.Next = Next
}

func (n *Node) updatePrev(Prev *Node) {
	n.Previous = Prev
}
