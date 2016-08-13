package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostLogin(t *testing.T) {
	var tests = []struct {
		Body     string
		Response map[string]interface{}
		Status   int
	}{
		{
			"uahruhh.r'hrjhauaeua",
			map[string]interface{}{
				"error": "Not Found",
			},
			400,
		},
	}

	for _, test := range tests {
		r, _ := http.NewRequest("POST", "/login", nil)
		w := httptest.NewRecorder()

		LoginHandler(w, r)

		if w.Code != test.Status {
			t.Errorf("Expected status %d got %v", test.Status, w.Code)
		}
	}
}
