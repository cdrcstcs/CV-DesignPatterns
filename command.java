// public Certainly! The Command design pattern is used to encapsulate a request as an object, thereby allowing for parameterization of clients with different requests, queuing of requests, and logging of the requests. Here's a complex example of the Command pattern in Java:

// Copy
import java.util.ArrayList;
import java.util.List;

// Command interface
interface Command {
    void execute();
}

// ConcreteCommand classes
class LightOnCommand implements Command {
    private Light light;

    public LightOnCommand(Light light) {
        this.light = light;
    }

    @Override
    public void execute() {
        light.turnOn();
    }
}

class LightOffCommand implements Command {
    private Light light;

    public LightOffCommand(Light light) {
        this.light = light;
    }

    @Override
    public void execute() {
        light.turnOff();
    }
}

// Receiver class
class Light {
    public void turnOn() {
        System.out.println("Light is ON");
    }

    public void turnOff() {
        System.out.println("Light is OFF");
    }
}

// Invoker class
class RemoteControl {
    private List<Command> commandHistory = new ArrayList<>();

    public void pressButton(Command command) {
        command.execute();
        commandHistory.add(command);
    }

    public void showCommandHistory() {
        System.out.println("Command History:");
        for (Command command : commandHistory) {
            System.out.println(command.getClass().getSimpleName());
        }
    }
}

// Client code
public class CommandPatternExample {
    public static void main(String[] args) {
        // Creating receiver
        Light light = new Light();

        // Creating commands
        Command lightOnCommand = new LightOnCommand(light);
        Command lightOffCommand = new LightOffCommand(light);

        // Creating invoker
        RemoteControl remoteControl = new RemoteControl();

        // Pressing buttons and executing commands
        remoteControl.pressButton(lightOnCommand);
        remoteControl.pressButton(lightOffCommand);

        // Showing command history
        remoteControl.showCommandHistory();
    }
}
// In this example:

// Command is the command interface with the execute method.
// LightOnCommand and LightOffCommand are concrete command classes implementing the Command interface for turning the light on and off, respectively.
// Light is the receiver class that performs the actual operations.
// RemoteControl is the invoker class that holds and executes commands, and also maintains a command history.
// This example models a remote control that can turn a light on and off. Commands are encapsulated as objects, and the invoker (RemoteControl) uses these commands without knowing the details of the operations. The command history is also maintained, allowing for undo/redo functionality or other use cases. {
    
