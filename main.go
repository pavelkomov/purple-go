package main

import (
	"fmt"
	"net/http"
)

func main() {

	router := http.NewServeMux()
	NewHelloHandler(router)
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on 8081")
	server.ListenAndServe()

}

func getCode() {

	resp, err := http.Get("https://ya.ru")
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(resp.StatusCode)

}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")

}
