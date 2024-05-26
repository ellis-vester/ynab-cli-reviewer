package tui

import (
	"github.com/brunomvsouza/ynab.go/api/budget"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/ellis-vester/ynab-cli-reviewer/ynab"
	"github.com/thoas/go-funk"
)

type ReviewModel struct {
	Client ynab.Ynab

	Budgets []*budget.Summary

	BudgetForm      huh.Form
	TransactionForm huh.Form
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
				Options(budgetOptions...),
		),
	)

	return &ReviewModel{
		Client:     client,
		BudgetForm: *budgetForm,
	}
}

func (m ReviewModel) Init() tea.Cmd {
	return m.BudgetForm.Init()
}

func (m ReviewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m ReviewModel) View() string {
	return m.BudgetForm.View()
}
