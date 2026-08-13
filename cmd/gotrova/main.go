package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gafreax/gotrova/internal/tui"
	"github.com/gafreax/gotrova/pkg/goapi"
)

func main() {
	baseURL := os.Getenv("GOTROVA_API_URL")
	client := goapi.NewClient(baseURL)
	model := tui.New(client)

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting Gotrova: %v\n", err)
		os.Exit(1)
	}
}
