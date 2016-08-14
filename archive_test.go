package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Logs in to the application and return the JWT
func logIn(w *httptest.ResponseRecorder) string {
	var response Response

	MyUser = User{Name: "McTest", Password: "ATestPassword"}

	r, _ := http.NewRequest(
		"POST",
		"/login",
		strings.NewReader("{\"Name\": \"McTest\", \"Password\": \"ATestPassword\"}"),
	)

	r.Header.Set("Content-Type", "application/json")

	LoginHandler(w, r)

	if w.Code != http.StatusOK {
		return ""
	}

	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		return ""
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return ""
	}

	token, ok := response["token"]
	if !ok {
		return ""
	}

	return token.(string)
}

func TestListArchiveHandler(t *testing.T) {
	r, _ := http.NewRequest("GET", "/archives", nil)
	w := httptest.NewRecorder()

	ListArchiveHandler(w, r)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d but got %v", http.StatusUnauthorized, w.Code)
	}
}

func TestListAddArchiveHandler(t *testing.T) {
	r, _ := http.NewRequest("GET", "/archives", nil)
	w := httptest.NewRecorder()

	logIn(w)
	ListArchiveHandler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %v", http.StatusOK, w.Code)
	}
}
