package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/yanmoyy/backend-assignment/internal/api"
	"github.com/yanmoyy/backend-assignment/internal/db"
)

func HandlerReset(w http.ResponseWriter, r *http.Request) {
	db.ClearIssues()
}

func HandlerCreateIssue(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var p api.CreateIssueParams
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
		user = db.GetUserByID(*p.UserId)
		if user == nil { // not found
			respondWithError(w, http.StatusBadRequest, "User not found", nil)
			return
		}
	}
	issue := api.Issue{
		ID:          db.GetNextIssueID(),
		Title:       p.Title,
		Description: p.Description,
		Status:      status,
		User:        user,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	db.AddIssue(issue)
	respondWithJSON(w, http.StatusCreated, issue)
}

func HandlerGetIssuesList(w http.ResponseWriter, r *http.Request) {
	// filter by status
	q := r.URL.Query()
	status := q.Get("status")
	var newIssues []api.Issue
	if status == "" {
		newIssues = db.GetIssuesAll()
	} else {
		newIssues = db.GetIssuesFilteredByStatus(status)
	}
	respondWithJSON(w, http.StatusOK, api.GetIssuesResponse{
		Issues: newIssues,
	})
}

func HandelrGetIssue(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid issue ID", err)
		return
	}
	issue, ok := db.GetIssueByID(id)
	if !ok {
		respondWithError(w, http.StatusNotFound, "Issue not found", nil)
		return
	}
	respondWithJSON(w, http.StatusOK, issue)
}

func HandlerUpdateIssue(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid issue ID", err)
		return
	}
	var p api.UpdateIssueParams
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&p)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	issue, ok := db.GetIssueByID(id)
	if !ok {
		respondWithError(w, http.StatusNotFound, "Issue not found", nil)
		return
	}
	issue, err = updateIssue(p, issue)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}
	db.SetIssueByID(id, issue)
	respondWithJSON(w, http.StatusOK, issue)
}

func updateIssue(p api.UpdateIssueParams, issue api.Issue) (api.Issue, error) {
	if issue.Status == api.StatusCompleted || issue.Status == api.StatusCancelled {
		return issue, fmt.Errorf("cannot update current issue")
	}
	if p.Title != "" {
		issue.Title = p.Title
	}
	if p.Description != "" {
		issue.Description = p.Description
	}
	if p.Status != "" {
		if p.UserId == nil && p.Status != api.StatusPending && p.Status != api.StatusCancelled {
			return issue, fmt.Errorf("need to specify userId")
		}
		issue.Status = p.Status
	} else {
		if issue.Status == api.StatusPending && p.UserId != nil {
			issue.Status = api.StatusInProgress
		} else if issue.User != nil && p.UserId == nil {
			issue.Status = api.StatusPending
		}
	}
	if p.UserId != nil {
		issue.User = db.GetUserByID(*p.UserId)
	}
	issue.UpdatedAt = time.Now()
	return issue, nil
}
