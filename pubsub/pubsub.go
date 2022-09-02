package main

import (
	"fmt"
)

// in memory implementation of a message queue
type MyQueue struct {
	Data        []string
	Subscribers map[string]int
}

func main() {

	var queue MyQueue

	queue.Subscribers = make(map[string]int)
	queue.Publish("Test1")
	fmt.Println(queue)
	queue.Publish("Test2")
	fmt.Println(queue)
	message := queue.Subscribe("subs1")
	fmt.Println(message)
	message = queue.Subscribe("subs2")
	fmt.Println(message)
	message = queue.Subscribe("subs1")
	fmt.Println(message)
}

// Publish publishes a given message to the queue
func (q *MyQueue) Publish(message string) {
	q.Data = append(q.Data, message)
}

// Subscribe returns first message from the queue and update the queue
func (q *MyQueue) Subscribe(subsid string) string {
	var message string
	if _, ok := q.Subscribers[subsid]; !ok {
		message = q.Data[0]
		q.Subscribers[subsid] = 1
	} else {
		if len(q.Data) > q.Subscribers[subsid] {
			message = q.Data[q.Subscribers[subsid]]
			q.Subscribers[subsid] += 1
		}
	}

	return message
}
