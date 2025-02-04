package factory

import (
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Error("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)
	if msg != "10.30 paid using cash\n" {
		t.Errorf("The cash payment method message was not correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Error("A payment method of type DebitCard must exist")
	}

	msg := payment.Pay(22.30)
	if msg != "22.30 paid using debit card\n" {
		t.Errorf("The debit card payment method message was not correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	payment, err := GetPaymentMethod(20)
	if payment != nil {
		t.Error("A payment method with ID 20 must return nil")
	}
	if err == nil {
		t.Error("An error must be returned when the payment method is not found")
	}
	t.Log("LOG:", err)
}
