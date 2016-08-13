package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Response map[string]interface{}

func TestPostLogin(t *testing.T) {
	var tests = []struct {
		Body     string
		Response Response
		Status   int
	}{
		{
			"uahruhh.r'hrjhauaeua",
			Response{
				"error": "Must send content as text/json",
			},
			400,
		},
	}

	for _, test := range tests {
		var response Response

		r, _ := http.NewRequest("POST", "/login", nil)
		w := httptest.NewRecorder()

		LoginHandler(w, r)

		if w.Code != test.Status {
			t.Errorf("Expected status %d got %v", test.Status, w.Code)
		}

		b, _ := ioutil.ReadAll(w.Body)
		json.Unmarshal(b, &response)

		for key, v := range test.Response {
			if _, ok := response[key]; !ok {
				t.Errorf("Expected %s:%v but key is missing in %v", key, v, len(response))
			} else if response[key] != v {
				t.Errorf("Expected %s:%v got %v", key, v, response[key])
			}
		}

	}
}
