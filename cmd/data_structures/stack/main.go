package main

import (
	"fmt"
	"go_tutorial/data_structures/stack"
	"log"
)

func main() {
	log.Default().Println("Initializing queue")
	s := stack.NewStack()
	s.ShowStack()

	fmt.Println()
	log.Default().Println("Pushing elements")
	s.Push(0)
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	s.Push(50)
	s.ShowStack()

	fmt.Println()
	log.Default().Println("Removing element")
	s.Pop()
	s.ShowStack()

	fmt.Println()
	log.Default().Println("Peeking stack")
	n := s.Peek()
	if n != nil {
		fmt.Println(n.Val)
	}

	fmt.Println()
	log.Default().Println("Size: ")
	fmt.Print(s.Size())
}
