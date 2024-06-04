package main

import (
	"fmt"
	"log"
)

// PaymentMethod is the interface
type PaymentMethod interface {
	Pay(amount float64) error
}

// CreditCard is a method that can process a payment
type CreditCard struct{}

// Pay for CreditCard
func (c CreditCard) Pay(amount float64) error {
	fmt.Printf("Paid %.2f using Credit Card\n", amount)
	return nil
}

// PayPal is another method that can process a payment
type PayPal struct{}

// Pay for PayPal
func (p PayPal) Pay(amount float64) error {
	fmt.Printf("Paid %.2f using PayPal\n", amount)
	return nil
}

// Bitcoin is another method that can process a payment
type Bitcoin struct{}

// Pay for Bitcoin
func (b Bitcoin) Pay(amount float64) error {
	fmt.Printf("Paid %.2f using Bitcoin\n", amount)
	return fmt.Errorf("no wallet selected")
}

// main function
func main() {
	// Use CreditCard as our payment method
	paymentMethod := CreditCard{}

	err := paymentMethod.Pay(300.20)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
