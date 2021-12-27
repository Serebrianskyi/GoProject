package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Email string `json:"email"`
	RegDate string `json:"regdate"`
}

var persons []Person

func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	personListBytes, err := json.Marshal(persons)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(personListBytes)
}

func createPersonHandler(w http.ResponseWriter, r *http.Request) {

	person := Person{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	person.Name = r.Form.Get("name")
	person.Email = r.Form.Get("email")
	person.RegDate = r.Form.Get("regdate")

	persons = append(persons, person)

	http.Redirect(w, r, "/assets/", http.StatusFound)
}