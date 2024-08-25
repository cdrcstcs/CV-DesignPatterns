// The Decorator design pattern is used to add new functionality to an object dynamically without altering its structure. Here's a complex example of the Decorator pattern in Java:

// Copy
// Component interface
interface Coffee {
    double cost();
    String description();
}

// ConcreteComponent class
class SimpleCoffee implements Coffee {
    @Override
    public double cost() {
        return 2.0; // Base cost of a simple coffee
    }

    @Override
    public String description() {
        return "Simple Coffee";
    }
}

// Decorator abstract class
abstract class CoffeeDecorator implements Coffee {
    protected Coffee decoratedCoffee;

    public CoffeeDecorator(Coffee decoratedCoffee) {
        this.decoratedCoffee = decoratedCoffee;
    }

    @Override
    public double cost() {
        return decoratedCoffee.cost();
    }

    @Override
    public String description() {
        return decoratedCoffee.description();
    }
}

// ConcreteDecorator classes
class MilkDecorator extends CoffeeDecorator {
    public MilkDecorator(Coffee decoratedCoffee) {
        super(decoratedCoffee);
    }

    @Override
    public double cost() {
        return super.cost() + 1.0; // Adding cost of milk
    }

    @Override
    public String description() {
        return super.description() + " with Milk";
    }
}

class SugarDecorator extends CoffeeDecorator {
    public SugarDecorator(Coffee decoratedCoffee) {
        super(decoratedCoffee);
    }

    @Override
    public double cost() {
        return super.cost() + 0.5; // Adding cost of sugar
    }

    @Override
    public String description() {
        return super.description() + " with Sugar";
    }
}

class VanillaDecorator extends CoffeeDecorator {
    public VanillaDecorator(Coffee decoratedCoffee) {
        super(decoratedCoffee);
    }

    @Override
    public double cost() {
        return super.cost() + 1.5; // Adding cost of vanilla
    }

    @Override
    public String description() {
        return super.description() + " with Vanilla";
    }
}

// Client code
public class DecoratorPatternExample {
    public static void main(String[] args) {
        Coffee simpleCoffee = new SimpleCoffee();
        System.out.println("Cost: $" + simpleCoffee.cost() + ", Description: " + simpleCoffee.description());

        Coffee milkCoffee = new MilkDecorator(simpleCoffee);
        System.out.println("Cost: $" + milkCoffee.cost() + ", Description: " + milkCoffee.description());

        Coffee sugarMilkCoffee = new SugarDecorator(milkCoffee);
        System.out.println("Cost: $" + sugarMilkCoffee.cost() + ", Description: " + sugarMilkCoffee.description());

        Coffee vanillaSugarMilkCoffee = new VanillaDecorator(sugarMilkCoffee);
        System.out.println("Cost: $" + vanillaSugarMilkCoffee.cost() + ", Description: " + vanillaSugarMilkCoffee.description());
    }
}
// In this example:

// The Coffee interface defines the component that we want to decorate.
// SimpleCoffee is a concrete component that implements the Coffee interface.
// CoffeeDecorator is an abstract class that implements the Coffee interface and holds a reference to a Coffee object.
// MilkDecorator, SugarDecorator, and VanillaDecorator are concrete decorator classes that extend CoffeeDecorator and add specific functionality (cost and description) to the decorated coffee.
// The client code demonstrates how different decorators can be combined to create a variety of coffee options. The decorators add additional behavior (cost and description) to the base SimpleCoffee object dynamically.