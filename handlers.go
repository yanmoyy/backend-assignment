package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// hardcoded users for assignment
var users = []User{
	{ID: 1, Name: "김개발"},
	{ID: 2, Name: "이디자인"},
	{ID: 3, Name: "박기획"},
}

func getUserByID(id uint) *User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

var issues = []Issue{}

func getNextIssueID() uint {
	return uint(len(issues) + 1) // #nosec G115
}

type CreateIssueParmas struct {
	Title       string `json:"title"`       // required
	Description string `json:"description"` // optional
	UserId      *uint  `json:"userId"`      // optional
}

func handlerCreateIssue(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create issue")
	decoder := json.NewDecoder(r.Body)
	var p CreateIssueParmas
	err := decoder.Decode(&p)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	if p.Title == "" {
		respondWithError(w, http.StatusBadRequest, "Title is required", nil)
		return
	}
	status := StatusPending
	var user *User
	if p.UserId != nil {
		status = StatusInProgress
		user = getUserByID(*p.UserId)
		if user == nil { // not found
			respondWithError(w, http.StatusBadRequest, "User not found", nil)
			return
		}
	}
	issue := Issue{
		ID:          getNextIssueID(),
		Title:       p.Title,
		Description: p.Description,
		Status:      status,
		User:        user,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	issues = append(issues, issue)
	respondWithJSON(w, http.StatusCreated, issue)
}

func handlerGetIssuesList(w http.ResponseWriter, r *http.Request) {
}

func handelrGetIssue(w http.ResponseWriter, r *http.Request) {
}
func handlerUpdateIssue(w http.ResponseWriter, r *http.Request) {
}
