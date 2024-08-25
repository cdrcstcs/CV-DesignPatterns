// Certainly! Here's an example of the Observer design pattern implemented in Go:

// Copy
package main

import (
	"fmt"
	"sync"
)

// Subject interface defines methods for adding, removing, and notifying observers
type Subject interface {
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// Observer interface defines the update method
type Observer interface {
	Update(subject Subject)
}

// ConcreteSubject struct is a concrete implementation of the Subject interface
type ConcreteSubject struct {
	state     int
	observers []Observer
	mutex     sync.Mutex
}

// AddObserver adds an observer to the list
func (s *ConcreteSubject) AddObserver(observer Observer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.observers = append(s.observers, observer)
}

// RemoveObserver removes an observer from the list
func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers notifies all observers of a change
func (s *ConcreteSubject) NotifyObservers() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, observer := range s.observers {
		observer.Update(s)
	}
}

// SetState sets the state of the subject and notifies observers
func (s *ConcreteSubject) SetState(state int) {
	s.state = state
	s.NotifyObservers()
}

// ConcreteObserver struct is a concrete implementation of the Observer interface
type ConcreteObserver struct {
	name string
}

// Update implements the update method to react to changes in the subject's state
func (o *ConcreteObserver) Update(subject Subject) {
	if concreteSubject, ok := subject.(*ConcreteSubject); ok {
		newState := concreteSubject.state
		fmt.Printf("Observer %s has been notified. New state: %d\n", o.name, newState)
	}
}

func main() {
	subject := &ConcreteSubject{}

	observer1 := &ConcreteObserver{name: "Observer 1"}
	observer2 := &ConcreteObserver{name: "Observer 2"}

	// Adding observers to the subject
	subject.AddObserver(observer1)
	subject.AddObserver(observer2)

	// Changing the state of the subject will notify the observers
	subject.SetState(10)

	// Removing observer1
	subject.RemoveObserver(observer1)

	// Changing the state again will notify only observer2
	subject.SetState(20)
}
main()
// In this example:

// ConcreteSubject is the concrete implementation of the Subject interface.
// ConcreteObserver is the concrete implementation of the Observer interface.
// The main function demonstrates how to use the Observer pattern in Go by creating a subject, adding observers, and changing the subject's state to see how observers are notified.
// This example uses a mutex to synchronize access to the observers' list, ensuring thread safety when adding or removing observers.