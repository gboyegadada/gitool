package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)


type model struct {
	choices  []string    
	cursor   int   
	selected map[int]struct{} 
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

			switch msg.String() {

			case "ctrl+c", "q":
					return m, tea.Quit

			case "up", "k":
					if m.cursor > 0 {
							m.cursor--
					}

			case "down", "j":
					if m.cursor < len(m.choices)-1 {
							m.cursor++
					}

			case "enter", " ":
					_, ok := m.selected[m.cursor]
					if ok {
							delete(m.selected, m.cursor)
					} else {
							m.selected[m.cursor] = struct{}{}
					}
			}
	}

	return m, nil
}

func (m model) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	for i, choice := range m.choices {

			cursor := " " // no cursor
			if m.cursor == i {
					cursor = ">" // cursor!
			}

			checked := " " 
			if _, ok := m.selected[i]; ok {
					checked = "x"
			}

			// Render the row
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	return s
}

func initialModel() model {
	return model{
		choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		selected: make(map[int]struct{}),
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
	}
}