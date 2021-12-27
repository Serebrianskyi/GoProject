package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetPersonsHandler(t *testing.T) {

	persons = []Person{
		{"Qwerty", "qwe@rty", "12.12.12"},
	}

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(getPersonHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := Person{"Qwerty", "qwe@rty", "12.12.12"}
	p := []Person{}
	err = json.NewDecoder(recorder.Body).Decode(&p)

	if err != nil {
		t.Fatal(err)
	}

	actual := p[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
func TestCreatePersonsHandler(t *testing.T) {

	persons = []Person{
		{"Qwerty", "qwe@rty", "12.12.12"},
	}

	form := newCreatePersonForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(createPersonHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := Person{"Qwerty2", "qwe2@rty", "12.12.13"}

	if err != nil {
		t.Fatal(err)
	}

	actual := persons[1]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func newCreatePersonForm() *url.Values {
	form := url.Values{}
	form.Set("name", "Qwerty2")
	form.Set("email", "qwe2@rty")
	form.Set("regdate", "12.12.13")
	return &form
}
