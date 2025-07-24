package queue

import "go_tutorial/data_structures/linkedlist"

type Queue struct {
	//O primeiro node da linkedlist será o primeiro item da fila
	//FIFO
	ll *linkedlist.LinkedList
}

func NewQueue() *Queue {
	return &Queue{
		ll: linkedlist.NewLinkedList(),
	}
}

func (q *Queue) Enqueue(value int) {
	n := linkedlist.NewNode(value)
	q.ll.Insert(n)
}

func (q *Queue) Dequeue() {
	q.ll.RemoveFirst()
}

func (q *Queue) Peek() *linkedlist.Node {
	return q.ll.Get(0)
}

func (q *Queue) ShowQueue() {
	q.ll.ShowList()
}

func (q *Queue) Size() int {
	return q.ll.Lenght
}
