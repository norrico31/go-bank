package bnk

import (
	"fmt"
	"time"
)

type Customer struct {
	name     string
	accounts []BankAccount
	AuditInfo
}

func (c *Customer) AddAccount(account BankAccount) {
	c.accounts = append(c.accounts, account)
}

func (c Customer) DisplayAccounts() {
	fmt.Printf("Customer: %s\n", c.name)
	for _, acct := range c.accounts {
		acct.DisplayBalance()
	}
}

func NewCustomer(name string) *Customer {
	return &Customer{
		name:      name,
		AuditInfo: AuditInfo{createdAt: time.Now(), updatedAt: time.Now(), deletedAt: time.Now(), lastModified: time.Now()},
	}
}
