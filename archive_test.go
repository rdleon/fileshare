package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListArchiveHandler(t *testing.T) {
	r, _ := http.NewRequest("GET", "/archives", nil)
	w := httptest.NewRecorder()

	ListArchiveHandler(w, r)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d but got %v", http.StatusUnauthorized, w.Code)
	}
}
