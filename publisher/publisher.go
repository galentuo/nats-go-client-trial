package main

import (
	"fmt"

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

	me := &person{Name: "derek", Age: 4, Address: "140 New Montgomery Street, San Francisco, CA"}

	// Go type Publisher
	c.Publish("hello", me)

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

	// // Close connection
	// c.Close()
}
