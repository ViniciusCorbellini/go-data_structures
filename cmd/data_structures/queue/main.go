package main

import (
	"fmt"
	"go_tutorial/data_structures/queue"
	"log"
)

func main() {
	fmt.Println()
	log.Default().Println("Initializing queue")
	q := queue.NewQueue()
	q.ShowQueue()

	fmt.Println()
	log.Default().Println("Inserting elements")
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.ShowQueue()

	fmt.Println()
	log.Default().Println("Dequeueing")
	q.Dequeue()
	q.ShowQueue()

	fmt.Println()
	log.Default().Println("Peeking queue")
	n := q.Peek()
	if n != nil {
		fmt.Println(n.Val)
	}

	fmt.Println()
	log.Default().Println("Size: ")
	fmt.Println(q.Size())
}
