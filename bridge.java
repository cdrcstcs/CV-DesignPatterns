// The Bridge design pattern is used to separate abstraction from implementation, allowing them to vary independently. Here's a complex example in Java where we'll model the abstraction of a Shape hierarchy and its drawing implementation.

// Copy
// Abstraction
interface Shape {
    void draw();
}

// Implementor
interface DrawingAPI {
    void drawCircle(double x, double y, double radius);

    void drawRectangle(double x, double y, double width, double height);
}

// Concrete Implementor
class DrawingAPI1 implements DrawingAPI {
    @Override
    public void drawCircle(double x, double y, double radius) {
        System.out.printf("API1 - Drawing Circle at (%.2f, %.2f) with radius %.2f%n", x, y, radius);
    }

    @Override
    public void drawRectangle(double x, double y, double width, double height) {
        System.out.printf("API1 - Drawing Rectangle at (%.2f, %.2f) with width %.2f and height %.2f%n", x, y, width, height);
    }
}

class DrawingAPI2 implements DrawingAPI {
    @Override
    public void drawCircle(double x, double y, double radius) {
        System.out.printf("API2 - Drawing Circle at (%.2f, %.2f) with radius %.2f%n", x, y, radius);
    }

    @Override
    public void drawRectangle(double x, double y, double width, double height) {
        System.out.printf("API2 - Drawing Rectangle at (%.2f, %.2f) with width %.2f and height %.2f%n", x, y, width, height);
    }
}

// Refined Abstraction
class Circle implements Shape {
    private double x, y, radius;
    private DrawingAPI drawingAPI;

    public Circle(double x, double y, double radius, DrawingAPI drawingAPI) {
        this.x = x;
        this.y = y;
        this.radius = radius;
        this.drawingAPI = drawingAPI;
    }

    @Override
    public void draw() {
        drawingAPI.drawCircle(x, y, radius);
    }
}

class Rectangle implements Shape {
    private double x, y, width, height;
    private DrawingAPI drawingAPI;

    public Rectangle(double x, double y, double width, double height, DrawingAPI drawingAPI) {
        this.x = x;
        this.y = y;
        this.width = width;
        this.height = height;
        this.drawingAPI = drawingAPI;
    }

    @Override
    public void draw() {
        drawingAPI.drawRectangle(x, y, width, height);
    }
}

public class BridgePatternExample {
    public static void main(String[] args) {
        DrawingAPI api1 = new DrawingAPI1();
        DrawingAPI api2 = new DrawingAPI2();

        Shape circle1 = new Circle(1, 2, 3, api1);
        Shape circle2 = new Circle(5, 7, 11, api2);

        Shape rectangle1 = new Rectangle(2, 3, 6, 9, api1);
        Shape rectangle2 = new Rectangle(4, 6, 8, 12, api2);

        circle1.draw();
        circle2.draw();
        rectangle1.draw();
        rectangle2.draw();
    }
}
In this example:

Shape is the abstraction that can be either a circle or a rectangle.
DrawingAPI is the interface for the drawing implementation.
DrawingAPI1 and DrawingAPI2 are concrete implementations of DrawingAPI.
Circle and Rectangle are refined abstractions that use a specific DrawingAPI.
The client code (BridgePatternExample) demonstrates how different shapes can use different drawing APIs.
This example illustrates how the Bridge pattern allows you to decouple the abstraction (Shape) from its implementation (DrawingAPI), making it easier to extend and maintain the code.