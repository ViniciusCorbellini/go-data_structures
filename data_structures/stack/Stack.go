package stack

import "go_tutorial/data_structures/linkedlist"

//O ultimo node a ser inserido será o primeiro a ser retirado
//LIFO
type Stack struct {
	ll *linkedlist.LinkedList
}

func NewStack() *Stack {
	return &Stack{
		ll: linkedlist.NewLinkedList(),
	}
}

func (s *Stack) Push(val int) {
	n := linkedlist.NewNode(val)
	s.ll.Insert(n)
}

func (s *Stack) Pop() {
	s.ll.RemoveLast()
}

func (s *Stack) Peek() *linkedlist.Node {
	last_index := s.ll.Lenght - 1
	return s.ll.Get(last_index)
}

func (s *Stack) ShowStack() {
	s.ll.ShowList()
}

func (s *Stack) Size() int {
	return s.ll.Lenght
}
