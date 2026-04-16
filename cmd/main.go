package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-quit:
		if sig == syscall.SIGINT {
			fmt.Print("Shutdown...")
		}
	case err := <-serverError:
		panic(err)
	}

	ctx, close := context.WithTimeout(context.Background(), time.Second * 30)
	defer close()

	server.Shutdown(ctx)
}