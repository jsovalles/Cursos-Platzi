package design_patterns

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using Cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using Bank account %d\n", bankAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func AdapterExample() {
	cash := &CashPayment{}
	ProcessPayment(cash)
	// bank := &BankPayment{}
	// ProcessPayment(bank)
	bpa := &BankPaymentAdapter{
		bankAccount: 5,
		BankPayment: &BankPayment{},
	}
	ProcessPayment(bpa)
}
