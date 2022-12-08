package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CatFacts struct {
	Id         string `json:"id"`
	Text       string `json:"text"`
	Source     string `json:"source"`
	Source_url string `json:"source_url"`
	Language   string `json:"language"`
	Permalink  string `json:"permalink"`
}

func response(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("https://uselessfacts.jsph.pl/random.json")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var catfacts CatFacts

	json.NewDecoder(resp.Body).Decode(&catfacts)

	fmt.Fprintf(w, catfacts.Text)
}

func main() {
	http.HandleFunc("/", response)
	http.ListenAndServe(":3000", nil)
}
