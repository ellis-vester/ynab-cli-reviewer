package tui

import (
	"strconv"
	"time"

	"github.com/brunomvsouza/ynab.go/api/budget"
	"github.com/brunomvsouza/ynab.go/api/transaction"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/ellis-vester/ynab-cli-reviewer/ynab"
	"github.com/thoas/go-funk"
)

type ReviewModel struct {
	Client ynab.Ynab

	Budgets      []*budget.Summary
	Transactions []*transaction.Transaction

	TransactionsFetched bool

	BudgetForm      *huh.Form
	TransactionForm *huh.Form
}

func NewReviewModel(client ynab.Ynab) *ReviewModel {

	budgets, err := client.GetBudgets()
	if err != nil {
		panic(err)
	}

	budgetOptions := funk.Map(budgets, func(x *budget.Summary) huh.Option[string] {
		return huh.NewOption(x.Name, x.ID)
	}).([]huh.Option[string])

	budgetForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select Budget").
				Options(budgetOptions...).
				Key("budget"),
		),
	)

	return &ReviewModel{
		Client:     client,
		BudgetForm: budgetForm,
	}
}

func (m ReviewModel) Init() tea.Cmd {
	return m.BudgetForm.Init()
}

func (m ReviewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case getTransactionsMsg:
		if msg.Err != nil {
			panic(msg.Err)
		}

		m.TransactionsFetched = true

		m.Transactions = msg.Transactions

		transOpts := funk.Map(msg.Transactions, func(x *transaction.Transaction) huh.Option[string] {
			return huh.NewOption(*x.CategoryName+" - "+strconv.Itoa(int(x.Amount)), x.ID)
		}).([]huh.Option[string])

		m.TransactionForm = huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					Title("Categorize Transactions").
					Options(transOpts...),
			),
		)
	}

	budgetForm, cmd := m.BudgetForm.Update(msg)
	if f, ok := budgetForm.(*huh.Form); ok {
		m.BudgetForm = f
		cmds = append(cmds, cmd)
	}

	if m.TransactionForm != nil {
		transForm, cmd := m.TransactionForm.Update(msg)
		if f, ok := transForm.(*huh.Form); ok {
			m.TransactionForm = f
			cmds = append(cmds, cmd)
		}
	}

	if m.BudgetForm.State == huh.StateCompleted && !m.TransactionsFetched {
		cmds = append(cmds, getTransactions(m))
	}

	return m, tea.Batch(cmds...)
}

func (m ReviewModel) View() string {

	if m.BudgetForm.State == huh.StateCompleted && len(m.Transactions) != 0 {
		return m.TransactionForm.View()
	} else {
		return m.BudgetForm.View()
	}
}

// Messages
type getTransactionsMsg struct {
	Err          error
	Transactions []*transaction.Transaction
}

func getTransactions(m ReviewModel) tea.Cmd {
	return func() tea.Msg {

		time.Sleep(1 * time.Second)

		budgetId := m.BudgetForm.GetString("budget")

		transactions, err := m.Client.GetPendingTransactions(budgetId)
		if err != nil {
			return getTransactionsMsg{
				Err:          err,
				Transactions: nil,
			}
		}

		return getTransactionsMsg{
			Err:          nil,
			Transactions: transactions,
		}
	}
}
