package goapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("q") != "fmt" {
			t.Errorf("expected query param 'q' to be 'fmt', got '%s'", r.URL.Query().Get("q"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"packages": [{"name": "fmt", "path": "fmt", "synopsis": "Package fmt implements formatted I/O."}], "count": 1}`))
	}))
	defer server.Close()

	client := NewClient(server.URL)
	res, err := client.Search(context.Background(), "fmt")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res.Count != 1 {
		t.Errorf("expected count 1, got %d", res.Count)
	}
	if len(res.Packages) != 1 || res.Packages[0].Name != "fmt" {
		t.Errorf("unexpected packages: %+v", res.Packages)
	}
}

func TestSearchMalformedJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"packages": [{"name": "fmt"`)) // malformed
	}))
	defer server.Close()

	client := NewClient(server.URL)
	_, err := client.Search(context.Background(), "fmt")
	if err == nil {
		t.Fatalf("expected error due to malformed json, got nil")
	}
}

func TestGetPackageDetailsSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name": "fmt", "path": "fmt", "synopsis": "fmt package"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL)
	res, err := client.GetPackageDetails(context.Background(), "fmt")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res.Name != "fmt" {
		t.Errorf("expected name 'fmt', got %s", res.Name)
	}
}

func TestGetPackageDetailsMalformedJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name": "fmt",`)) // malformed
	}))
	defer server.Close()

	client := NewClient(server.URL)
	_, err := client.GetPackageDetails(context.Background(), "fmt")
	if err == nil {
		t.Fatalf("expected error due to malformed json, got nil")
	}
}
