package main

import (
	"errors"
	"fmt"
	"sync"
)

type stack struct {
	m      sync.Mutex
	values []string
}

func (s *stack) pop() (string, error) {
	s.m.Lock()
	defer s.m.Unlock()
	if len(s.values) == 0 {
		return "", errors.New("no values in stack")
	}
	valueToPop := s.values[len(s.values)-1]
	s.values = (s.values)[:len(s.values)-1]
	return valueToPop, nil
}

func (s *stack) push(value string) {
	s.m.Lock()
	defer s.m.Unlock()
	s.values = append(s.values, value)
}

func main() {
	var myStack stack
	myStack.push("item 1")
	myStack.push("item 2")
	myStack.push("item 3")
	myStack.push("item 4")
	fmt.Printf("stack values %v\n", myStack.values)
	poppedVal, err := myStack.pop()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("popped %v\n", poppedVal)
	fmt.Printf("stack values %v\n", myStack.values)
	poppedVal, err = myStack.pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("poppedVal %v\n", poppedVal)
	fmt.Printf("stack values %v\n", myStack.values)
}
