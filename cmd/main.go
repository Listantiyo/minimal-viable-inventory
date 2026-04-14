package main

import (
	"fmt"
	"net/http"
	// "syscall"
	// "errors"
)

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	serverError := make(chan error, 1)
	go func() {
		serverError <- server.ListenAndServe()
	}()

	err := <-serverError
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}