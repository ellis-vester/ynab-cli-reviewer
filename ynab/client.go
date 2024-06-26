package ynab

import (
	"github.com/brunomvsouza/ynab.go"
	"github.com/brunomvsouza/ynab.go/api/budget"
	"github.com/brunomvsouza/ynab.go/api/transaction"
)

type Ynab struct {
	Client ynab.ClientServicer
}

func (c *Ynab) GetBudgets() ([]*budget.Summary, error) {

	budgets, err := c.Client.Budget().GetBudgets()
	if err != nil {
		return nil, err
	}

	return budgets, nil
}

func (c *Ynab) GetPendingTransactions(budgetId string) ([]*transaction.Transaction, error) {

	transactions, err := c.Client.Transaction().GetTransactions(budgetId, nil)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
