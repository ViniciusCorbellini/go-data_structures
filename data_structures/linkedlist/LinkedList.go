package linkedlist

import (
	"fmt"
)

type LinkedList struct {
	Start_ptr *Node
	End       *Node
	Lenght    int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Start_ptr: &Node{},
		Lenght:    0,
		End:       nil,
	}
}

func (ll *LinkedList) Insert(n *Node) {
	if ll.IsEmpty() {
		ll.Start_ptr.updateNext(n)
		ll.End = n
	} else {
		ll.End.updateNext(n)
		n.updatePrev(ll.End)
		ll.End = n
	}
	ll.Lenght++
}

func (ll *LinkedList) Remove(index int) {
	if ll.IsEmpty() || index < 0 || index >= ll.Lenght {
		return
	}

	if index == 0 {
		ll.RemoveFirst()
		return
	}

	if index >= ll.Lenght-1 {
		ll.RemoveLast()
		return
	}

	curr := ll.Start_ptr
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	curr.updateNext(curr.Next.Next)
	ll.Lenght--
}

func (ll *LinkedList) RemoveFirst() {
	if ll.IsEmpty() {
		return
	}

	if ll.Lenght == 1 {
		ll.removeIfLenOne()
	} else {
		ll.Start_ptr.Next = ll.Start_ptr.Next.Next
		ll.Start_ptr.Next.updatePrev(ll.Start_ptr)
	}
	ll.Lenght--
}

func (ll *LinkedList) RemoveLast() {
	if ll.IsEmpty() {
		return
	}

	if ll.Lenght == 1 {
		ll.removeIfLenOne()
	} else {
		ll.End = ll.End.Previous
		ll.End.updateNext(nil)
	}
	ll.Lenght--
}

func (ll *LinkedList) RemoveAllOCurrencesOf(val int) {
	for ll.Start_ptr.Next != nil && ll.Start_ptr.Next.Val == val {
		ll.RemoveFirst()
	}

	for ll.End != nil && ll.End.Val == val {
		ll.RemoveLast()
	}

	curr := ll.Start_ptr.Next
	for curr != nil && curr.Next != nil {
		if curr.Next.Val == val {
			toRemove := curr.Next
			curr.Next = toRemove.Next
			if curr.Next != nil {
				curr.Next.updatePrev(curr)
			}
			ll.Lenght--
		} else {
			curr = curr.Next
			if curr.Next != nil {
				curr.Next.updatePrev(curr)
			}
		}
	}
}

func (ll *LinkedList) Reverse() {
	if ll.Lenght <= 1 {
		return
	}

	ll.Start_ptr.Next.updatePrev(nil) //Tira a ligacao do Nó inicial com o ponteiro sentinela
	ll.Start_ptr.updateNext(ll.End)   //Liga o sentinela ao fim da LL

	curr := ll.End
	for prev := curr.Previous; prev != nil; curr = curr.Next {
		curr.Next = prev
		mid := prev
		prev = prev.Previous
		mid.updatePrev(curr)
		mid.updateNext(prev)
	}
	ll.End = curr                              //Atualiza o fim da lista
	ll.Start_ptr.Next.updatePrev(ll.Start_ptr) //Encadeia duplamente o comeco com o sentinela
}

func (ll *LinkedList) removeIfLenOne() {
	ll.Start_ptr.Next = nil
	ll.End = nil
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.Lenght == 0
}

func (ll *LinkedList) ShowList() {
	fmt.Printf("START -> ")
	for curr := ll.Start_ptr.Next; curr != nil; curr = curr.Next {
		fmt.Printf("%v -> ", curr.Val)
	}
	fmt.Print("NULL\n")
	fmt.Printf("Length: %v\n", ll.Lenght)

	if ll.Lenght == 0 {
		return
	}
	fmt.Printf("Starting Node: %v // Ending Node: %v\n", ll.Start_ptr.Next.Val, ll.End.Val)
}
