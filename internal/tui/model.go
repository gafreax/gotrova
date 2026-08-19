package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gafreax/gotrova/pkg/goapi"
)

type state int

const (
	stateInput state = iota
	stateLoading
	stateList
	stateError
)

type item struct {
	pkg goapi.Package
}

func (i item) Title() string {
	if i.pkg.Path != "" {
		return i.pkg.Path
	}
	return i.pkg.Name
}

func (i item) Description() string {
	desc := i.pkg.Synopsis
	if desc == "" {
		desc = "No synopsis available."
	}
	if i.pkg.Version != "" {
		return fmt.Sprintf("[%s] %s", i.pkg.Version, desc)
	}
	return desc
}

func (i item) FilterValue() string {
	return i.pkg.Path + " " + i.pkg.Name + " " + i.pkg.Synopsis
}

type Model struct {
	client goapi.Client

	state    state
	input    textinput.Model
	spinner  spinner.Model
	list     list.Model
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

	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.Color("#06B6D4")).
		Foreground(lipgloss.Color("#06B6D4")).
		Bold(true).
		Padding(0, 0, 0, 1)

	delegate.Styles.SelectedDesc = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.Color("#06B6D4")).
		Foreground(lipgloss.Color("#94A3B8")).
		Padding(0, 0, 0, 1)

	delegate.Styles.NormalTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E2E8F0")).
		Padding(0, 0, 0, 2)

	delegate.Styles.NormalDesc = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#64748B")).
		Padding(0, 0, 0, 2)

	l := list.New([]list.Item{}, delegate, 0, 0)
	l.Title = "GOTROVA // Search Results"
	l.Styles.Title = lipgloss.NewStyle().
		Background(lipgloss.Color("#06B6D4")).
		Foreground(lipgloss.Color("#0F172A")).
		Bold(true).
		Padding(0, 1)

	return Model{
		client:  client,
		state:   stateInput,
		input:   ti,
		spinner: sp,
		list:    l,
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
