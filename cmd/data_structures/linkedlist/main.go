package main

import (
	"fmt"
	"log"
)

func main() {
	//Initializing the ll
	fmt.Println()
	log.Default().Println("\nInitializing the ll")
	ll := newLinkedList()
	ll.showList()

	//Inserting elements
	fmt.Println()
	log.Default().Println("Inserting Elements")
	ll.Insert(newNode(0))
	ll.Insert(newNode(10))
	ll.Insert(newNode(10))
	ll.Insert(newNode(20))
	ll.Insert(newNode(10))
	ll.Insert(newNode(30))
	ll.Insert(newNode(10))
	ll.Insert(newNode(40))
	ll.Insert(newNode(10))
	ll.showList()

	//Removing the last element
	fmt.Println()
	log.Default().Println("Removing the last element")
	ll.RemoveLast()
	ll.showList()

	//Removing the first element
	fmt.Println()
	log.Default().Println("Removing the first element")
	ll.RemoveFirst()
	ll.showList()

	//Removing the n-element
	fmt.Println()
	log.Default().Println("Removing the 2nd element")
	ll.Remove(1)
	ll.showList()

	//Removing all ocurrences of '10'
	fmt.Println()
	log.Default().Println("Removing all ocurrences of '10'")
	ll.RemoveAllOCurrencesOf(10)
	ll.showList()

	//Reversing
	fmt.Println()
	log.Default().Println("Reversing linked list")
	ll.Reverse()
	ll.showList()
}
