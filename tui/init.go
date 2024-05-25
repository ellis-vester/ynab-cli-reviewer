package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/ellis-vester/ynab-cli-reviewer/config"
	"github.com/spf13/viper"
)

type InitModel struct {
	Form   *huh.Form
	ApiKey string
}

func NewInitModel() *InitModel {
	model := InitModel{}

	model.Form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("YNAB API Key").
				Description("Saved to ~/.ynabr/config.json for future use.").
				Key("api-key").
				Value(&model.ApiKey),
		),
	)

	return &model
}

func (m InitModel) Init() tea.Cmd {
	return m.Form.Init()
}

func (m InitModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		}
	case apiKeySavedMsg:
		if msg.Err != nil {
			panic(msg.Err)
		}

		return m, tea.Quit
	}

	var cmds []tea.Cmd

	form, cmd := m.Form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.Form = f
		cmds = append(cmds, cmd)
	}

	if m.Form.State == huh.StateCompleted {
		apiKey := m.Form.GetString("api-key")
		cmds = append(cmds, saveApiKeyToConfig(apiKey))
	}

	return m, tea.Batch(cmds...)
}

func (m InitModel) View() string {
	return m.Form.View()
}

// Messages
type apiKeySavedMsg struct {
	Err error
}

// Commands
func saveApiKeyToConfig(apiKey string) tea.Cmd {
	return func() tea.Msg {

		conf := config.Config{
			ApiKey: apiKey,
		}
		viper.Set("config", &conf)

		err := viper.WriteConfig()
		if err != nil {
			fmt.Errorf("Failed to write to viper config", err)
			return apiKeySavedMsg{
				Err: err,
			}
		}

		return apiKeySavedMsg{
			Err: nil,
		}
	}
}
