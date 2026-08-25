package goapi

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
)

// Client defines the interface for interacting with the pkg.go.dev API
type Client interface {
	Search(ctx context.Context, query string) (*SearchResult, error)
	GetPackageDetails(ctx context.Context, path string) (*Package, error)
}

type defaultClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new Client with the given base URL.
func NewClient(baseURL string) Client {
	if baseURL == "" {
		baseURL = "https://pkg.go.dev"
	}
	return &defaultClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *defaultClient) Search(ctx context.Context, query string) (*SearchResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	u, err := url.Parse(c.baseURL + "/search")
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	q := u.Query()
	q.Set("q", query)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}
	html := string(body)

	// Naive HTML parsing with regex
	var result SearchResult
	
	// Split by SearchSnippet
	snippets := strings.Split(html, "class=\"SearchSnippet\"")
	if len(snippets) > 1 {
		for _, snip := range snippets[1:] {
			var pkg Package
			
			// Extract title and path
			reTitle := regexp.MustCompile(`data-test-id="snippet-title"[^>]*>\s*([^\n<]+)\s*<span class="SearchSnippet-header-path">\(([^)]+)\)`)
			if matches := reTitle.FindStringSubmatch(snip); len(matches) >= 3 {
				pkg.Name = strings.TrimSpace(matches[1])
				pkg.Path = strings.TrimSpace(matches[2])
			} else {
				continue // skip if we can't even find a title
			}
			
			// Extract synopsis
			reSyn := regexp.MustCompile(`data-test-id="snippet-synopsis"[^>]*>\s*([^\n<]+)`)
			if matches := reSyn.FindStringSubmatch(snip); len(matches) >= 2 {
				pkg.Synopsis = strings.TrimSpace(matches[1])
			}
			
			// Extract version
			reVer := regexp.MustCompile(`<strong>(v[^<]+)</strong>\s*published`)
			if matches := reVer.FindStringSubmatch(snip); len(matches) >= 2 {
				pkg.Version = strings.TrimSpace(matches[1])
			}
			
			result.Packages = append(result.Packages, pkg)
		}
	}
	
	// Prioritize matches in package name or path
	lowerQuery := strings.ToLower(query)
	sort.SliceStable(result.Packages, func(i, j int) bool {
		pi := result.Packages[i]
		pj := result.Packages[j]
		
		iInName := strings.Contains(strings.ToLower(pi.Name), lowerQuery) || strings.Contains(strings.ToLower(pi.Path), lowerQuery)
		jInName := strings.Contains(strings.ToLower(pj.Name), lowerQuery) || strings.Contains(strings.ToLower(pj.Path), lowerQuery)
		
		if iInName && !jInName {
			return true
		}
		if !iInName && jInName {
			return false
		}
		return false
	})
	
	result.Count = len(result.Packages)
	return &result, nil
}

func (c *defaultClient) GetPackageDetails(ctx context.Context, path string) (*Package, error) {
	// Not implemented for this minimal TUI
	return nil, fmt.Errorf("not implemented")
}
