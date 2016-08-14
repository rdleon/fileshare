package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type Response map[string]interface{}

func TestInvalidHeaderLogin(t *testing.T) {
	var expected, response Response

	expected = Response{"error": "Must send content as application/json"}

	r, _ := http.NewRequest("POST", "/login", strings.NewReader("TeST"))
	w := httptest.NewRecorder()

	LoginHandler(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d got %v", http.StatusBadRequest, w.Code)
	}

	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error("Can't read the response body")
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		t.Error("Can't read the response json")
	}

	for key, v := range expected {
		if _, ok := response[key]; !ok {
			t.Errorf("Expected %s:%v but key is missing in %v", key, v, len(response))
		} else if response[key] != v {
			t.Errorf("Expected %s:%v got %v", key, v, response[key])
		}
	}

}

func TestMalformedJSONLogin(t *testing.T) {
	var expected, response Response

	expected = Response{"error": "Bad Request"}

	r, _ := http.NewRequest("POST", "/login", strings.NewReader("TeST"))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	LoginHandler(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d got %v", http.StatusBadRequest, w.Code)
	}

	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error("Can't read the response body")
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		t.Error("Can't read the response json")
	}

	for key, v := range expected {
		if _, ok := response[key]; !ok {
			t.Errorf("Expected %s:%v but key is missing in %v", key, v, len(response))
		} else if response[key] != v {
			t.Errorf("Expected %s:%v got %v", key, v, response[key])
		}
	}
}

func TestCorrectLogin(t *testing.T) {
	var response Response
	MyUser = User{Name: "McTest", Password: "ATestPassword"}

	r, _ := http.NewRequest("POST", "/login", strings.NewReader("{\"Name\": \"McTest\", \"Password\": \"ATestPassword\"}"))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	LoginHandler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d got %v", http.StatusOK, w.Code)
	}

	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error("Can't read the response body")
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		t.Error("Can't read the response json")
	}

	if _, ok := response["token"]; !ok {
		t.Errorf("Expected token but it is missing in %v", response)
	}

	r, _ = http.NewRequest("POST", "/login", strings.NewReader("{\"Name\": \"rdleon\", \"Password\": \"password\"}"))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+response["token"].(string))
	w = httptest.NewRecorder()

	LoginHandler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d got %v", http.StatusOK, w.Code)
	}

	b, err = ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error("Can't read the response body")
	}

	response = make(Response)
	err = json.Unmarshal(b, &response)
	if err != nil {
		t.Error("Can't read the response json")
	}

	if _, ok := response["loggedin"]; !ok {
		t.Errorf("Expected token but it is missing in %v", response)
	}
}

func TestLogoutHandler(t *testing.T) {
	r, _ := http.NewRequest("GET", "/logout", nil)
	w := httptest.NewRecorder()

	LogoutHandler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status to be '%d' but got '%v'", http.StatusOK, w.Code)
	}
}
