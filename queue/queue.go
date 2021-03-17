package main

import (
	"errors"
	"fmt"
	"sync"
)

type queue struct {
	m      sync.Mutex
	values []string
}

func (q *queue) dequeue() (string, error) {
	q.m.Lock()
	defer q.m.Unlock()

	if len(q.values) == 0 {
		return "", errors.New("no values in queue")
	}

	dequeuedVal := q.values[0]
	q.values = q.values[1:]
	return dequeuedVal, nil
}

func (q *queue) enqueue(value string) {
	q.m.Lock()
	defer q.m.Unlock()
	q.values = append(q.values, value)
}

func main() {
	var myQueue queue
	myQueue.enqueue("item 1")
	myQueue.enqueue("item 2")
	myQueue.enqueue("item 3")
	myQueue.enqueue("item 4")
	fmt.Printf("queue values %v\n", myQueue.values)
	dequeuedVal, err := myQueue.dequeue()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("dequeued %v\n", dequeuedVal)
	fmt.Printf("queue values %v\n", myQueue.values)
	dequeuedVal, err = myQueue.dequeue()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("dequeued %v\n", dequeuedVal)
	fmt.Printf("queue values %v\n", myQueue.values)
}
