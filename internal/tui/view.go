package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)

	// Header styles
	bannerStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#06B6D4")).
			Foreground(lipgloss.Color("#0F172A")).
			Bold(true).
			Padding(0, 1)

	subHeaderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#64748B")).
			MarginLeft(1)

	headerBoxStyle = lipgloss.NewStyle().
			MarginBottom(1)

	// Component containers
	panelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#334155")).
			Padding(1, 2)

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#06B6D4")).
			Bold(true).
			MarginBottom(1)

	hintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#64748B")).
			MarginTop(1)

	// Status styles
	loadingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#06B6D4")).
			Bold(true)

	loadingQueryStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#F8FAFC")).
				Italic(true)

	// Error styles
	errorBadgeStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#EF4444")).
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true).
			Padding(0, 1)

	errorBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#EF4444")).
			Padding(1, 2).
			MarginTop(1)

	errorMsgStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FCA5A5"))
)

func renderHeader() string {
	badge := bannerStyle.Render("GOTROVA")
	sub := subHeaderStyle.Render("pkg.go.dev package index")
	return headerBoxStyle.Render(badge + sub)
}

func (m Model) View() string {
	switch m.state {
	case stateInput:
		header := renderHeader()
		label := labelStyle.Render("SEARCH PACKAGES")
		inputView := m.input.View()
		hint := hintStyle.Render("Press Enter to search • Esc to quit")

		content := fmt.Sprintf("%s\n\n%s\n%s", label, inputView, hint)
		return docStyle.Render(header + "\n" + panelStyle.Render(content))

	case stateLoading:
		header := renderHeader()
		status := fmt.Sprintf(
			"%s %s %s",
			m.spinner.View(),
			loadingStyle.Render("Searching pkg.go.dev for"),
			loadingQueryStyle.Render(fmt.Sprintf("%q...", m.input.Value())),
		)
		return docStyle.Render(header + "\n" + panelStyle.Render(status))

	case stateList:
		header := renderHeader()
		tableView := panelStyle.Render(m.table.View())
		hint := hintStyle.Render("Press Enter to view details • Esc to search again")
		return docStyle.Render(header + "\n" + tableView + "\n" + hint)

	case stateDetail:
		header := renderHeader()
		hint := hintStyle.Render("Press Esc to go back • ↑/↓ to scroll")
		return docStyle.Render(header + "\n" + m.viewport.View() + "\n" + hint)

	case stateError:
		header := renderHeader()
		errBadge := errorBadgeStyle.Render("ERROR")
		errText := errorMsgStyle.Render(m.errorMsg)
		hint := hintStyle.Render("Press Enter to try again • Esc to quit")

		content := fmt.Sprintf("%s\n\n%s\n\n%s", errBadge, errText, hint)
		return docStyle.Render(header + "\n" + errorBoxStyle.Render(content))
	}

	return ""
}
