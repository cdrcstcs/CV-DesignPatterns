// Certainly! Here's an example of the Decorator design pattern implemented in Go:

// Copy
package main

import "fmt"

// Component interface
type Coffee interface {
	Cost() float64
	Description() string
}

// ConcreteComponent struct
type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() float64 {
	return 2.0 // Base cost of a simple coffee
}

func (c *SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// Decorator abstract class
type CoffeeDecorator struct {
	coffee Coffee
}

func (d *CoffeeDecorator) Cost() float64 {
	return d.coffee.Cost()
}

func (d *CoffeeDecorator) Description() string {
	return d.coffee.Description()
}

// ConcreteDecorator classes
type MilkDecorator struct {
	*CoffeeDecorator
}

func NewMilkDecorator(coffee Coffee) *MilkDecorator {
	return &MilkDecorator{&CoffeeDecorator{coffee}}
}

func (d *MilkDecorator) Cost() float64 {
	return d.coffee.Cost() + 1.0 // Adding cost of milk
}

func (d *MilkDecorator) Description() string {
	return d.coffee.Description() + " with Milk"
}

type SugarDecorator struct {
	*CoffeeDecorator
}

func NewSugarDecorator(coffee Coffee) *SugarDecorator {
	return &SugarDecorator{&CoffeeDecorator{coffee}}
}

func (d *SugarDecorator) Cost() float64 {
	return d.coffee.Cost() + 0.5 // Adding cost of sugar
}

func (d *SugarDecorator) Description() string {
	return d.coffee.Description() + " with Sugar"
}

type VanillaDecorator struct {
	*CoffeeDecorator
}

func NewVanillaDecorator(coffee Coffee) *VanillaDecorator {
	return &VanillaDecorator{&CoffeeDecorator{coffee}}
}

func (d *VanillaDecorator) Cost() float64 {
	return d.coffee.Cost() + 1.5 // Adding cost of vanilla
}

func (d *VanillaDecorator) Description() string {
	return d.coffee.Description() + " with Vanilla"
}

// Client code
func main() {
	simpleCoffee := &SimpleCoffee{}
	printCoffee(simpleCoffee)

	milkCoffee := NewMilkDecorator(simpleCoffee)
	printCoffee(milkCoffee)

	sugarMilkCoffee := NewSugarDecorator(milkCoffee)
	printCoffee(sugarMilkCoffee)

	vanillaSugarMilkCoffee := NewVanillaDecorator(sugarMilkCoffee)
	printCoffee(vanillaSugarMilkCoffee)
}

func printCoffee(coffee Coffee) {
	fmt.Printf("Cost: $%.2f, Description: %s\n", coffee.Cost(), coffee.Description())
}
In this example:

Coffee is the component interface.
SimpleCoffee is the concrete component implementing the Coffee interface.
CoffeeDecorator is the decorator abstract class, embedding the Coffee interface to ensure it implements the Coffee methods.
MilkDecorator, SugarDecorator, and VanillaDecorator are concrete decorator classes that embed CoffeeDecorator and add specific functionality to the decorated coffee.
The client code demonstrates how different decorators can be combined to create various coffee options. The decorators add additional behavior (cost and description) to the base SimpleCoffee object dynamically.