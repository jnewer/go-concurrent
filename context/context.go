package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func simulateTask(ctx context.Context) {
	select {
	case <-time.After(8 * time.Second):
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Task  cancelled")

	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	go simulateTask(ctx)

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintln(w, "Request processed successfully")
	case <-ctx.Done():
		fmt.Fprintln(w, "Request timeout")
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
