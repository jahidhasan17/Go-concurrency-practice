package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type PubSub struct {
	subs 	map[string][]chan string
	mu 	 	sync.Mutex
	closed 	bool
}

func (pb *PubSub) Subscribe(topic string) <-chan string {
	pb.mu.Lock()
	defer pb.mu.Unlock()

	if pb.closed{
		return nil
	}

	var ch chan string = make(chan string)

	pb.subs[topic] = append(pb.subs[topic], ch)

	return ch
}

func (pb *PubSub) Publish(topic string, content string) {
	pb.mu.Lock()
	defer pb.mu.Unlock()

	if pb.closed{
		return
	}

	for _, sub := range pb.subs[topic] {
		sub <- content
	}
}

func NewPubSub() *PubSub {
	return &PubSub{
		subs: make(map[string][]chan string),
	}
}

func (pb *PubSub)Close()  {
	pb.mu.Lock()
	defer pb.mu.Unlock()

	if pb.closed {
		return
	}

	pb.closed = true

	for _, subs := range pb.subs{
		for _, sub := range subs{
			close(sub)
		}
	}

	pb.subs = nil
}

func main() {
	pb := NewPubSub()

	go func() {
		for i := 0; i < 3; i++ {
			sub := pb.Subscribe("news")
			go func(sub <-chan string, i int) {
				for msg := range sub{
					fmt.Println("Getting message from Subscriber", i, "and Message is", msg)
				}
			}(sub, i)
		}
	}()

	// This for Subscriber to be ready
	time.Sleep(time.Millisecond * 500)


	for j := 0; j < 1; j++ {
		go func (pub int)  {
			for i := 0; i < 2; i++{
				pb.Publish("news", "From Publisher " + strconv.Itoa(pub) + " Message " + strconv.Itoa(i))
				time.Sleep(time.Millisecond * 500)
			}
		}(j)
	}
	
	time.Sleep(time.Second * 3)
	pb.Close()
}