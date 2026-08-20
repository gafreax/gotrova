package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gafreax/gotrova/pkg/goapi"
)

type state int

const (
	stateInput state = iota
	stateLoading
	stateList
	stateDetail
	stateError
)

type Model struct {
	client goapi.Client

	state    state
	input    textinput.Model
	spinner  spinner.Model
	table    table.Model
	viewport viewport.Model
	errorMsg string
}

func New(client goapi.Client) Model {
	ti := textinput.New()
	ti.Placeholder = "e.g. lipgloss, bubbletea, gin-gonic/gin..."
	ti.Prompt = "❯ "
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#06B6D4")).Bold(true)
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F8FAFC"))
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#64748B"))
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#06B6D4")).Bold(true)

	columns := []table.Column{
		{Title: "Package", Width: 35},
		{Title: "Version", Width: 15},
		{Title: "Description", Width: 60},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)
	
	vp := viewport.New(80, 10)
	vp.Style = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#334155"))

	return Model{
		client:   client,
		state:    stateInput,
		input:    ti,
		spinner:  sp,
		table:    t,
		viewport: vp,
	}
}

// SetQuery sets the initial search query and transitions the state to loading
func (m *Model) SetQuery(query string) {
	if query != "" {
		m.input.SetValue(query)
		m.state = stateLoading
	}
}

func (m Model) Init() tea.Cmd {
	if m.state == stateLoading {
		return tea.Batch(textinput.Blink, m.spinner.Tick, m.searchCmd(m.input.Value()))
	}
	return textinput.Blink
}
