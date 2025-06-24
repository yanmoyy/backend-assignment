package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/yanmoyy/backend-assignment/internal/api"
	"github.com/yanmoyy/backend-assignment/internal/client"
)

func handlerReset(w http.ResponseWriter, r *http.Request) {
	clearIssues()
}

func handlerCreateIssue(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var p client.CreateIssueParams
	err := decoder.Decode(&p)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	if p.Title == "" {
		respondWithError(w, http.StatusBadRequest, "Title is required", nil)
		return
	}
	status := api.StatusPending
	var user *api.User
	if p.UserId != nil {
		status = api.StatusInProgress
		user = getUserByID(*p.UserId)
		if user == nil { // not found
			respondWithError(w, http.StatusBadRequest, "User not found", nil)
			return
		}
	}
	issue := api.Issue{
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
	var newIssues []api.Issue
	if status == "" {
		newIssues = issues
	} else {
		newIssues = filterIssuesByStatus(issues, status)
	}
	type response struct {
		Issues []api.Issue `json:"issues"`
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
