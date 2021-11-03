package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Message string

type Greeter struct {
	Message Message // <- adding a Message field
	Grumpy  bool
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func NewMessage(phrase string) Message {
	return Message(phrase)
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	e, err := InitializeEvent("jjjjjj")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}
