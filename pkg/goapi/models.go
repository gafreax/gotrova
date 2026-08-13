package goapi

// Package represents a Go package retrieved from pkg.go.dev
type Package struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	Synopsis   string `json:"synopsis"`
	Version    string `json:"version"`
	License    string `json:"license"`
	CommitTime string `json:"commit_time"`
}

// SearchResult represents the top-level response from a search query
type SearchResult struct {
	Packages []Package `json:"packages"`
	Count    int       `json:"count"`
}
