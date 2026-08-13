package goapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Client defines the interface for interacting with the pkg.go.dev API
type Client interface {
	// Search queries the pkg.go.dev API and returns a list of packages
	Search(ctx context.Context, query string) (*SearchResult, error)

	// GetPackageDetails retrieves detailed information about a specific package path
	GetPackageDetails(ctx context.Context, path string) (*Package, error)
}

type defaultClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new Client with the given base URL.
func NewClient(baseURL string) Client {
	if baseURL == "" {
		baseURL = "https://pkg.go.dev/api"
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

	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

func (c *defaultClient) GetPackageDetails(ctx context.Context, path string) (*Package, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	u, err := url.Parse(fmt.Sprintf("%s/packages/%s", c.baseURL, url.PathEscape(path)))
	if err != nil {
		return nil, fmt.Errorf("invalid url: %w", err)
	}

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

	var pkg Package
	if err := json.NewDecoder(resp.Body).Decode(&pkg); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &pkg, nil
}
