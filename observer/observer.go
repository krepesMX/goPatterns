package main

import (
	"fmt"
)

type Topic interface {
	register(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

type EmailClient struct {
	id string
}

func (ec *EmailClient) updateValue(value string) {
	fmt.Printf("Sending Email - %s available from client %s", value, ec.id)
}

func (ec *EmailClient) getId() string {
	return ec.id
}

func main() {
	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &EmailClient{
		id: "12ab",
	}
	secondObserver := &EmailClient{
		id: "34dv",
	}

	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)

	nvidiaItem.UpdateAvailable()
}
