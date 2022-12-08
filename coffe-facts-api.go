package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CoffeeFacts struct {
	File string `json:"file"`
}

func responseCofeeFacts(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("https://coffee.alexflipnote.dev/random.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var coffeeFacts CoffeeFacts

	json.NewDecoder(resp.Body).Decode(&coffeeFacts)
	fmt.Fprintf(w, coffeeFacts.File)

}
func main() {
	http.HandleFunc("/coffee", responseCofeeFacts)
	http.ListenAndServe(":3001", nil)
}
