package main

import (
	"fmt"
	"go_tutorial/data_structures/linkedlist"
	"log"
)

func main() {
	//Initializing the ll
	fmt.Println()
	log.Default().Println("\nInitializing the ll")
	ll := linkedlist.NewLinkedList()
	ll.ShowList()

	//Inserting elements
	fmt.Println()
	log.Default().Println("Inserting Elements")
	ll.Insert(linkedlist.NewNode(0))
	ll.Insert(linkedlist.NewNode(10))
	ll.Insert(linkedlist.NewNode(10))
	ll.Insert(linkedlist.NewNode(20))
	ll.Insert(linkedlist.NewNode(10))
	ll.Insert(linkedlist.NewNode(30))
	ll.Insert(linkedlist.NewNode(10))
	ll.Insert(linkedlist.NewNode(40))
	ll.Insert(linkedlist.NewNode(10))
	ll.ShowList()

	//Removing the last element
	fmt.Println()
	log.Default().Println("Removing the last element")
	ll.RemoveLast()
	ll.ShowList()

	//Removing the first element
	fmt.Println()
	log.Default().Println("Removing the first element")
	ll.RemoveFirst()
	ll.ShowList()

	//Removing the n-element
	fmt.Println()
	log.Default().Println("Removing the 2nd element")
	ll.Remove(1)
	ll.ShowList()

	//Removing all ocurrences of '10'
	fmt.Println()
	log.Default().Println("Removing all ocurrences of '10'")
	ll.RemoveAllOCurrencesOf(10)
	ll.ShowList()

	//Reversing
	fmt.Println()
	log.Default().Println("Reversing linked list")
	ll.Reverse()
	ll.ShowList()

	//Getting index 0
	fmt.Println()
	log.Default().Println("Getting first element")
	n := ll.Get(0)
	if n != nil {
		fmt.Println(n.Val)
	}
}
