package main

import "github.com/yanmoyy/backend-assignment/internal/api"

// hardcoded users db
var users = []api.User{
	{ID: 1, Name: "김개발"},
	{ID: 2, Name: "이디자인"},
	{ID: 3, Name: "박기획"},
}

func getUserByID(id uint) *api.User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

var issues = []api.Issue{}

func getNextIssueID() uint {
	return uint(len(issues) + 1) // #nosec G115
}

func getIssueByID(id uint) *api.Issue {
	for _, issue := range issues {
		if issue.ID == id {
			return &issue
		}
	}
	return nil
}

func clearIssues() {
	issues = []api.Issue{}
}

func filterIssuesByStatus(issues []api.Issue, status string) []api.Issue {
	filteredIssues := []api.Issue{}
	for _, issue := range issues {
		if issue.Status == status {
			filteredIssues = append(filteredIssues, issue)
		}
	}
	return filteredIssues
}
