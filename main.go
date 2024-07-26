package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string         // items on to-do list
	cursor   int              // which to-do list item cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
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
