package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// hardcoded users for assignment
// var users = []User{
// 	{ID: 1, Name: "김개발"},
// 	{ID: 2, Name: "이디자인"},
// 	{ID: 3, Name: "박기획"},
// }

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

	fmt.Println("Server is running on port 8080")
	log.Fatal(srv.ListenAndServe())
}
