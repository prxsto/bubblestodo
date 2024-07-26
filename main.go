package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	selected map[int]struct{}
	choices  []string
	cursor   int
}

func initialModel() model {
	return model{
		// our todo list is a grocery list
		choices: []string{"buy carrots", "buy celery", "buy kohlrabi"},

		// a map which indicates choices are selected. we're using
		// the map like a mathematical set. the keys refer to the
		// indices of the "choices" slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// just return 'nil', which means "no I/O right now, please"
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// is it a key press?
	case tea.KeyMsg:

		// cool, what was the actual key pressed?
		switch msg.String() {

		// exit program
		case "ctrl+c", "q":
			return m, tea.Quit

		// "up" or "k" move cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// "down" or "j" move cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// "space" and "return" toggle selected
		// state for item under cursor
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

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
