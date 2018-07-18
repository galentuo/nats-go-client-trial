package main

import (
	"fmt"
	"time"

	nats "github.com/nats-io/go-nats"
)

func main() {
	fmt.Println(nats.DefaultURL)
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	fmt.Println(">>>> NATS <<<<")

	// // Simple Async Subscriber
	// c.Subscribe("foo", func(s string) {
	// 	fmt.Printf("Received a message: %s\n", s)
	// })
	// // Simple Publisher
	// c.Publish("foo", "Hello World")

	// EncodedConn can Publish any raw Go type using the registered Encoder
	type person struct {
		Name    string
		Address string
		Age     int
	}

	// Go type Subscriber
	c.Subscribe("hello", func(p *person) {
		fmt.Printf("1. Sub 1 Received a person: %+v\n", p)
	})

	// Go type Subscriber
	c.Subscribe("hello", func(p *person) {
		fmt.Printf("2. Sub 2 Received a person: %+v\n", p)
	})

	// All subscriptions with the same queue name will form a queue group.
	// Each message will be delivered to only one subscriber per queue group,
	// using queuing semantics. You can have as many queue groups as you wish.
	// Normal subscribers will continue to work as expected.
	c.QueueSubscribe("hello", "group1", func(p *person) {
		fmt.Printf("3. Group-1 Sub 1 Received a person: %+v\n", p)
	})

	c.QueueSubscribe("hello", "group1", func(p *person) {
		fmt.Printf("4. Group-1 Sub 2 Received a person: %+v\n", p)
	})

	c.QueueSubscribe("hello", "group2", func(p *person) {
		fmt.Printf("5. Group-2 Sub 1 Received a person: %+v\n", p)
	})

	c.QueueSubscribe("hello", "group2", func(p *person) {
		fmt.Printf("6. Group-2 Sub 2 Received a person: %+v\n", p)
	})

	// me := &person{Name: "derek", Age: 22, Address: "140 New Montgomery Street, San Francisco, CA"}

	// // Go type Publisher
	// c.Publish("hello", me)

	// // Requests
	// var response string
	// err := c.Request("help", "help me", &response, 10*time.Millisecond)
	// if err != nil {
	// 	fmt.Printf("Request failed: %v\n", err)
	// }

	// // Replying
	// c.Subscribe("help", func(subj, reply string, msg string) {
	// 	c.Publish(reply, "I can help!")
	// })

	for {
		time.Sleep(10 * time.Second)
	}

	// // Close connection
	// c.Close()
}
