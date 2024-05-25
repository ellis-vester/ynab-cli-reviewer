package ynab

import (
	"github.com/brunomvsouza/ynab.go"
	"github.com/brunomvsouza/ynab.go/api/transaction"
)

func GetPendingTransactions() ([]transaction.Transaction, error) {
	c := ynab.NewClient("")

	transactions, err := c.Transaction().GetTransactions()
	if err != nil {
		return nil, err
	}

	return transactions, nil

}
