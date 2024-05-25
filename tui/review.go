package tui

import tea "github.com/charmbracelet/bubbletea"

type ReviewModel struct{}

func NewReviewModel() *ReviewModel {
	return &ReviewModel{}
}

func (m ReviewModel) Init() tea.Cmd {
	return nil
}

func (m ReviewModel) Update(tea.Msg) (tea.Model, tea.Cmd) {

	return m, nil
}

func (m ReviewModel) View() string {
	return ""
}
