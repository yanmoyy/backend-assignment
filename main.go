package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/yanmoyy/backend-assignment/internal/server"
)

func main() {
	srv := http.Server{
		Addr:        ":8080",
		ReadTimeout: 10 * time.Second,
	}
	// post request /issue
	http.HandleFunc("POST /issue", server.HandlerCreateIssue)
	http.HandleFunc("GET /issues", server.HandlerGetIssuesList)
	http.HandleFunc("GET /issue/{id}", server.HandelrGetIssue)
	http.HandleFunc("PATCH /issue/{id}", server.HandlerUpdateIssue)

	// reset
	http.HandleFunc("POST /reset", server.HandlerReset)

	fmt.Println("Server is running on port 8080")
	log.Fatal(srv.ListenAndServe())
}
