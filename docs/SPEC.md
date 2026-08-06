# gotrova Specification

## Overview
gotrova is a Terminal User Interface (TUI) application designed for searching and discovering Go packages. It leverages the pkg.go.dev API to fetch package information.

## TUI States
The application will have the following distinct states:
1. **Input State**: The initial state where the user enters a search query.
2. **Loading State**: Displayed while the application is making network requests to the pkg.go.dev API.
3. **List View State**: Displays the search results in a scrollable list, allowing the user to select a package for more details.
4. **Error Handling State**: Displays an error message and offers retry or back options if an API request fails or another error occurs.

## Data Models
```go
package goapi

// Package represents a Go package retrieved from pkg.go.dev
type Package struct {
    Name        string
    Path        string
    Synopsis    string
    Version     string
    License     string
    CommitTime  string
}

// SearchResult represents the top-level response from a search query
type SearchResult struct {
    Packages []Package
    Count    int
}
```

## HTTP Client Interface
```go
package goapi

import "context"

// Client defines the interface for interacting with the pkg.go.dev API
type Client interface {
    // Search queries the pkg.go.dev API and returns a list of packages
    Search(ctx context.Context, query string) (*SearchResult, error)
    
    // GetPackageDetails retrieves detailed information about a specific package path
    GetPackageDetails(ctx context.Context, path string) (*Package, error)
}
```
