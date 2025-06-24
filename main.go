package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	srv := http.Server{
		Addr:        ":8080",
		ReadTimeout: 10 * time.Second,
	}
	// post request /issue
	http.HandleFunc("POST /issue", handlerCreateIssue)
	http.HandleFunc("GET /issues", handlerGetIssuesList)
	http.HandleFunc("GET /issue:id", handelrGetIssue)
	http.HandleFunc("PATCH /issue:id", handlerUpdateIssue)

	// reset
	http.HandleFunc("POST /reset", handlerReset)

	fmt.Println("Server is running on port 8080")
	log.Fatal(srv.ListenAndServe())
}
