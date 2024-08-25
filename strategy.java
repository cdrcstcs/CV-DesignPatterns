// The Strategy Design Pattern is a behavioral design pattern that defines a family of algorithms, encapsulates each algorithm, and makes them interchangeable. It allows a client to choose an algorithm from a family of algorithms at runtime.

// Here's a complex example in Java to illustrate the Strategy Design Pattern. Let's consider a scenario where we have a payment system that supports multiple payment methods (e.g., credit card, PayPal, Bitcoin), and we want to allow clients to choose their preferred payment method dynamically.

// Copy
// PaymentStrategy interface
public interface PaymentStrategy {
    void pay(int amount);
}

// Concrete implementation of PaymentStrategy for Credit Card
public class CreditCardPayment implements PaymentStrategy {
    private String cardNumber;

    public CreditCardPayment(String cardNumber) {
        this.cardNumber = cardNumber;
    }

    @Override
    public void pay(int amount) {
        // Perform credit card payment logic
        System.out.println("Paid " + amount + " using Credit Card: " + cardNumber);
    }
}

// Concrete implementation of PaymentStrategy for PayPal
public class PayPalPayment implements PaymentStrategy {
    private String email;

    public PayPalPayment(String email) {
        this.email = email;
    }

    @Override
    public void pay(int amount) {
        // Perform PayPal payment logic
        System.out.println("Paid " + amount + " using PayPal: " + email);
    }
}

// Context class that uses the PaymentStrategy
public class ShoppingCart {
    private PaymentStrategy paymentStrategy;

    public void setPaymentStrategy(PaymentStrategy paymentStrategy) {
        this.paymentStrategy = paymentStrategy;
    }

    public void checkout(int amount) {
        // Perform checkout logic
        System.out.println("Checkout initiated...");
        
        // Use the selected payment strategy to pay
        paymentStrategy.pay(amount);
    }
}

// Client code
public class StrategyPatternExample {
    public static void main(String[] args) {
        // Create payment strategies
        PaymentStrategy creditCardPayment = new CreditCardPayment("1234-5678-9101-1121");
        PaymentStrategy payPalPayment = new PayPalPayment("user@example.com");

        // Create a shopping cart
        ShoppingCart cart = new ShoppingCart();

        // Perform checkout with Credit Card payment
        cart.setPaymentStrategy(creditCardPayment);
        cart.checkout(100);

        // Perform checkout with PayPal payment
        cart.setPaymentStrategy(payPalPayment);
        cart.checkout(50);
    }
}
// In this example:

// PaymentStrategy is an interface that declares the pay method.
// CreditCardPayment and PayPalPayment are concrete implementations of the PaymentStrategy interface, representing different payment methods.
// ShoppingCart is a context class that uses the selected payment strategy during the checkout process.
// The client code creates instances of payment strategies and sets them in the shopping cart, allowing dynamic switching of payment methods at runtime.
// This example demonstrates how the Strategy Design Pattern enables flexibility by allowing the client to choose different algorithms (payment strategies) dynamically.