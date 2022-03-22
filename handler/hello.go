package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var people []string

type Salute struct {
	Country string
	Salute  string
}

type Hello struct {
	l *log.Logger
}

var names []string

func (h *Hello) SayHello(r http.ResponseWriter, req *http.Request) {
	user := chi.URLParam(req, "name")
	people = append(people, user)
	salute := randomSalutation()
	log.Println("Greetings in", salute.Country)
	greet := fmt.Sprintln(salute.Salute, user)
	r.Write([]byte(greet))
}

func (h *Hello) ListPeople(r http.ResponseWriter, req *http.Request) {
	r.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(people)
	r.Write(json)
}

func randomSalutation() Salute {
	salutations := readJSON()
	country := keys(salutations)
	randomCountry := country[rand.Intn(len(salutations))]
	salute := Salute{
		Country: randomCountry,
		Salute:  salutations[randomCountry],
	}

	return salute
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{
		l,
	}
}

func keys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func readJSON() map[string]string {
	content, err := ioutil.ReadFile("./data/hi.json")
	if err != nil {
		log.Fatal("Error when opening file", err)
	}

	var hi map[string]string
	err = json.Unmarshal(content, &hi)
	if err != nil {
		log.Println("Cannot read json file. Defaulting to hi", err)
		hi["default"] = "hi"
	}
	return hi
}
