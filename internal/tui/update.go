package tui

import (
	"context"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gafreax/gotrova/pkg/goapi"
)

type searchMsg struct {
	result *goapi.SearchResult
}

type errMsg struct {
	err error
}

func (m Model) searchCmd(query string) tea.Cmd {
	return func() tea.Msg {
		res, err := m.client.Search(context.Background(), query)
		if err != nil {
			return errMsg{err}
		}
		return searchMsg{result: res}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		_, v := docStyle.GetFrameSize()
		m.table.SetHeight(msg.Height - v - 4) // adjust for header and padding
	}

	switch m.state {
	case stateInput:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEsc:
				return m, tea.Quit
			case tea.KeyEnter:
				if m.input.Value() != "" {
					m.state = stateLoading
					return m, tea.Batch(
						m.spinner.Tick,
						m.searchCmd(m.input.Value()),
					)
				}
			}
		}
		m.input, cmd = m.input.Update(msg)
		cmds = append(cmds, cmd)

	case stateLoading:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEsc {
				return m, tea.Quit
			}
		case searchMsg:
			m.state = stateList
			var rows []table.Row
			for _, pkg := range msg.result.Packages {
				path := pkg.Path
				if path == "" {
					path = pkg.Name
				}
				rows = append(rows, table.Row{
					path,
					pkg.Version,
					pkg.Synopsis,
				})
			}
			m.table.SetRows(rows)
			m.table.GotoTop()
		case errMsg:
			m.state = stateError
			m.errorMsg = msg.err.Error()
		default:
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
		}

	case stateList:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEsc {
				m.state = stateInput
				m.input.SetValue("")
				m.input.Focus()
				return m, textinput.Blink
			}
		}
		m.table, cmd = m.table.Update(msg)
		cmds = append(cmds, cmd)

	case stateError:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEsc:
				return m, tea.Quit
			case tea.KeyEnter:
				m.state = stateInput
				m.input.SetValue("")
				m.input.Focus()
				return m, textinput.Blink
			}
		}
	}

	return m, tea.Batch(cmds...)
}
