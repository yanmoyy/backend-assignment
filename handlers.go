package main

import (
	"encoding/json"
	"net/http"
	"strconv"
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

func getIssueByID(id uint) *Issue {
	for _, issue := range issues {
		if issue.ID == id {
			return &issue
		}
	}
	return nil
}

func filterIssuesByStatus(issues []Issue, status string) []Issue {
	filteredIssues := []Issue{}
	for _, issue := range issues {
		if issue.Status == status {
			filteredIssues = append(filteredIssues, issue)
		}
	}
	return filteredIssues
}

type CreateIssueParmas struct {
	Title       string `json:"title"`       // required
	Description string `json:"description"` // optional
	UserId      *uint  `json:"userId"`      // optional
}

func handlerCreateIssue(w http.ResponseWriter, r *http.Request) {
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
	// filter by status
	q := r.URL.Query()
	status := q.Get("status")
	var newIssues []Issue
	if status == "" {
		newIssues = issues
	} else {
		newIssues = filterIssuesByStatus(issues, status)
	}
	type response struct {
		Issues []Issue `json:"issues"`
	}
	respondWithJSON(w, http.StatusOK, response{
		Issues: newIssues,
	})
}

func handelrGetIssue(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid issue ID", err)
		return
	}
	issue := getIssueByID(uint(id))
	if issue == nil {
		respondWithError(w, http.StatusNotFound, "Issue not found", nil)
		return
	}
	respondWithJSON(w, http.StatusOK, issue)
}

func handlerUpdateIssue(w http.ResponseWriter, r *http.Request) {
	// idString := r.PathValue("id")
	// id, err := strconv.ParseUint(idString, 10, 64)
	// if err != nil {
	// 	respondWithError(w, http.StatusBadRequest, "Invalid issue ID", err)
	// 	return
	// }
	// issue := getIssueByID(uint(id))
	// if issue == nil {
	// 	respondWithError(w, http.StatusNotFound, "Issue not found", nil)
	// 	return
	// }
	// decoder := json.NewDecoder(r.Body)
	// var p UpdateIssueParmas
	// err = decoder.Decode(&p)
	// if err != nil {
	// 	respondWithError(w, http.StatusBadRequest, "Invalid request body", err)
	// 	return
	// }
	// if p.Title != "" {
	// 	issue.Title = p.Title
	// }
	// if p.Description != "" {
	// 	issue.Description = p.Description
	// }
	// if p.Status != "" {
	// 	issue.Status = p.Status
	// }
	// if p.UserId != nil {
	// 	issue.User = getUserByID(*p.UserId)
	// }
	// respondWithJSON(w, http.StatusOK, issue)
}
