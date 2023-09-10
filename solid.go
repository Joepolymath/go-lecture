package main

import "fmt"

type PaymentMethod interface {
	Pay()
}

type Payment struct {}

// =============== Payment method section. Add payment methods below =====================
type CreditCard struct {
	amount	float64
}

type Paypal struct {
	amount	float64
}

// End of payment method section.

func (cc CreditCard) Pay() {
	fmt.Printf("Paid %.2f with Credit Card\n", cc.amount)
}

func (cc Paypal) Pay() {
	fmt.Printf("Paid %.2f with Credit Card\n", cc.amount)
}

func (p Payment) Process(pm PaymentMethod) {
	pm.Pay()
}

func Ocp() {
	payment := Payment{}
	method := CreditCard{amount: 200.0}
	paypal := Paypal{amount: 100.00}
	payment.Process(method)
	payment.Process(paypal)
}