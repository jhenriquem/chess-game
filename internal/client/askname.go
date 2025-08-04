package client

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func AskName() string {
	p := tea.NewProgram(initialModel())
	var _ tea.Model

	var baseModel tea.Model
	var err error
	if baseModel, err = p.Run(); err != nil {
		log.Fatal(err)
	}

	finalModel := baseModel.(Model)
	name := finalModel.Input

	return name
}

type (
	errMsg error
)

type Model struct {
	Input     string
	textInput textinput.Model
	err       error
}

func initialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Model{
		textInput: ti,
		err:       nil,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.Input = m.textInput.Value()
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return fmt.Sprintf(
		"♟️ Welcome to the Game \n What’s your name ?\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
