package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	RandomInteger int `json:"random_integer"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		res := Response{
			RandomInteger: rand.Intn(100) + 1, // Random integer between 1 and 100
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	http.ListenAndServe(":8080", nil)
}
