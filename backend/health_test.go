package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	if contentType := rec.Header().Get("Content-Type"); contentType != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", contentType)
	}

	var response HealthResponse

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	if response.Status != "healthy" {
		t.Errorf("expected status 'healthy', got %q", response.Status)
	}

	if response.Service != "technical-spec-api" {
		t.Errorf("expected service 'technical-spec-api', got %q", response.Service)
	}

	if response.Version != "0.1.0" {
		t.Errorf("expected version '0.1.0', got %q", response.Version)
	}
}
