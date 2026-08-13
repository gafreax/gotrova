package tui

import (
	"context"
	"errors"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gafreax/gotrova/pkg/goapi"
)

type mockClient struct {
	searchFunc func(ctx context.Context, query string) (*goapi.SearchResult, error)
}

func (m *mockClient) Search(ctx context.Context, query string) (*goapi.SearchResult, error) {
	if m.searchFunc != nil {
		return m.searchFunc(ctx, query)
	}
	return &goapi.SearchResult{}, nil
}

func (m *mockClient) GetPackageDetails(ctx context.Context, path string) (*goapi.Package, error) {
	return nil, nil
}

func TestNewModel(t *testing.T) {
	client := &mockClient{}
	m := New(client)

	if m.state != stateInput {
		t.Errorf("expected initial state to be stateInput, got %v", m.state)
	}

	view := m.View()
	if !strings.Contains(view, "GOTROVA") {
		t.Errorf("expected view to contain header GOTROVA, got %s", view)
	}
	if !strings.Contains(view, "SEARCH PACKAGES") {
		t.Errorf("expected view to contain SEARCH PACKAGES prompt, got %s", view)
	}
}

func TestModelSearchStateTransition(t *testing.T) {
	client := &mockClient{
		searchFunc: func(ctx context.Context, query string) (*goapi.SearchResult, error) {
			return &goapi.SearchResult{
				Count: 1,
				Packages: []goapi.Package{
					{
						Name:     "lipgloss",
						Path:     "github.com/charmbracelet/lipgloss",
						Synopsis: "Style definitions for nice terminal layouts",
						Version:  "v1.0.0",
					},
				},
			}, nil
		},
	}

	m := New(client)
	m.input.SetValue("lipgloss")

	// Simulate pressing enter
	enterMsg := tea.KeyMsg{Type: tea.KeyEnter}
	updatedModel, cmd := m.Update(enterMsg)
	m = updatedModel.(Model)

	if m.state != stateLoading {
		t.Fatalf("expected state to transition to stateLoading, got %v", m.state)
	}
	if cmd == nil {
		t.Fatalf("expected search cmd batch, got nil")
	}

	// Simulate successful search result message
	res, _ := client.Search(context.Background(), "lipgloss")
	updatedModel, _ = m.Update(searchMsg{result: res})
	m = updatedModel.(Model)

	if m.state != stateList {
		t.Fatalf("expected state to transition to stateList, got %v", m.state)
	}

	view := m.View()
	if !strings.Contains(view, "GOTROVA // Search Results") {
		t.Errorf("expected list view to contain list title, got %s", view)
	}
}

func TestModelErrorStateTransition(t *testing.T) {
	client := &mockClient{}
	m := New(client)

	// Simulate error msg
	updatedModel, _ := m.Update(errMsg{err: errors.New("network timeout")})
	m = updatedModel.(Model)

	if m.state != stateError {
		t.Fatalf("expected state to transition to stateError, got %v", m.state)
	}

	view := m.View()
	if !strings.Contains(view, "network timeout") {
		t.Errorf("expected view to contain error message, got %s", view)
	}

	// Press Enter to reset
	enterMsg := tea.KeyMsg{Type: tea.KeyEnter}
	updatedModel, _ = m.Update(enterMsg)
	m = updatedModel.(Model)

	if m.state != stateInput {
		t.Errorf("expected state to return to stateInput after pressing enter on error, got %v", m.state)
	}
}
