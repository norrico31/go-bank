package bnk

import "fmt"

type Bank struct {
	name      string
	customers []Customer
}

func (b *Bank) AddCustomer(customer Customer) {
	b.customers = append(b.customers, customer)
}

func (b Bank) DisplayCustomers() {
	fmt.Printf("Bank: %s\n", b.name)
	for _, customer := range b.customers {
		customer.DisplayAccounts()
	}
}

func NewBank(name string) *Bank {
	return &Bank{name: name}
}
