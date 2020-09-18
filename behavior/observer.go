package behavior

import "fmt"

type Event struct {
	id string
}

type Observer interface {
	Notify(event Event)
}

type observer struct {
	name string
}

func NewObserver(name string) Observer {
	return &observer{name}
}

func (o *observer) Notify(event Event) {
	fmt.Println(o.name, event.id)
}

type Notifier interface {
	Add(obs Observer)
	Remove(obs Observer)
	Notify(event Event)
}

type eventNotifier struct {
	observers []Observer
}

func NewEventNotifier() Notifier {
	return &eventNotifier{}
}

func (e *eventNotifier) Add(obs Observer) {
	e.observers = append(e.observers, obs)
}

func (e *eventNotifier) Remove(obs Observer) {
	for i := 0; i < len(e.observers); i++ {
		if obs == e.observers[i] {
			e.observers = append(e.observers[:i], e.observers[i+1:]...)
		}
	}
}

func (e *eventNotifier) Notify(event Event) {
	for i := 0; i < len(e.observers); i++ {
		e.observers[i].Notify(event)
	}
}
