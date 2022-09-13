package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct {
}

func (CashPayment) Pay() {
	fmt.Println("Payment using Cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct {
}

func (BankPayment) Pay(bancAccount int) {

	fmt.Printf("Payment using banck account %d", bancAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bancAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bancAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	bpa := &BankPaymentAdapter{
		bancAccount: 5,
		BankPayment: &BankPayment{},
	}
	ProcessPayment(bpa)

	str := "abc" + "123"
	
	fmt.Println(str)
}
