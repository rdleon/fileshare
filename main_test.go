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

	expected = Response{"error": "Must send content as text/json"}

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
