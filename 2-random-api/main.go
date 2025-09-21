package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

type Response struct {
	Number int `json:"number"`
}

func main() {

	http.HandleFunc("/roll", func(w http.ResponseWriter, r *http.Request) {

		num := rand.Intn(6) + 1

		w.Write([]byte(strconv.Itoa(num)))

	})

	http.ListenAndServe(":8081", nil)
}
