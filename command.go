// Certainly! Here's a complex example of the Command design pattern in Go:

// Copy
package main

import "fmt"

// Command interface
type Command interface {
	Execute()
}

// Receiver interface
type Light interface {
	TurnOn()
	TurnOff()
}

// ConcreteReceiver
type LivingRoomLight struct{}

func (l *LivingRoomLight) TurnOn() {
	fmt.Println("Living room light is ON")
}

func (l *LivingRoomLight) TurnOff() {
	fmt.Println("Living room light is OFF")
}

// ConcreteCommand
type TurnOnCommand struct {
	light Light
}

func NewTurnOnCommand(light Light) *TurnOnCommand {
	return &TurnOnCommand{light: light}
}

func (c *TurnOnCommand) Execute() {
	c.light.TurnOn()
}

// ConcreteCommand
type TurnOffCommand struct {
	light Light
}

func NewTurnOffCommand(light Light) *TurnOffCommand {
	return &TurnOffCommand{light: light}
}

func (c *TurnOffCommand) Execute() {
	c.light.TurnOff()
}

// Invoker
type RemoteControl struct {
	command Command
}

func (rc *RemoteControl) SetCommand(command Command) {
	rc.command = command
}

func (rc *RemoteControl) PressButton() {
	rc.command.Execute()
}

func main() {
	// Receiver
	livingRoomLight := &LivingRoomLight{}

	// Concrete Commands
	turnOnCommand := NewTurnOnCommand(livingRoomLight)
	turnOffCommand := NewTurnOffCommand(livingRoomLight)

	// Invoker
	remoteControl := &RemoteControl{}

	// Pressing buttons and executing commands
	remoteControl.SetCommand(turnOnCommand)
	remoteControl.PressButton()

	remoteControl.SetCommand(turnOffCommand)
	remoteControl.PressButton()
}
In this example:

Command is the command interface with the Execute method.
Light is the receiver interface with TurnOn and TurnOff methods.
LivingRoomLight is the concrete receiver.
TurnOnCommand and TurnOffCommand are concrete command classes implementing the Command interface.
RemoteControl is the invoker class.
The program demonstrates a remote control that can turn on and off a living room light. Commands are encapsulated as objects, and the invoker (RemoteControl) uses these commands without knowing the details of the operations. The use of interfaces allows for flexibility in adding new commands and receivers without modifying the invoker.