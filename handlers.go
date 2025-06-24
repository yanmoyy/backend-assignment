package main

import (
	"net/http"
)

func handlerCreateIssue(w http.ResponseWriter, r *http.Request) {
	println("create issue")
}

func handlerGetIssuesList(w http.ResponseWriter, r *http.Request) {
	println("get issue list")
}

func handelrGetIssue(w http.ResponseWriter, r *http.Request) {
	println("get issue")
}
func handlerUpdateIssue(w http.ResponseWriter, r *http.Request) {
	println("update issue")
}
