package db

import "github.com/yanmoyy/backend-assignment/internal/api"

// hardcoded db
var users = []api.User{
	{ID: 1, Name: "김개발"},
	{ID: 2, Name: "이디자인"},
	{ID: 3, Name: "박기획"},
}

var issues = []api.Issue{}

func GetUserByID(id uint) *api.User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

func GetNextIssueID() uint {
	return uint(len(issues) + 1) // #nosec G115
}

// #nosec G115
func GetIssueByID(id int) (api.Issue, bool) {
	for _, issue := range issues {
		if issue.ID == uint(id) {
			return issue, true
		}
	}
	return api.Issue{}, false
}

func AddIssue(issue api.Issue) {
	issues = append(issues, issue)
}

func GetIssuesAll() []api.Issue {
	return issues
}

func ClearIssues() {
	issues = []api.Issue{}
}

func GetIssuesFilteredByStatus(status string) []api.Issue {
	filteredIssues := []api.Issue{}
	for _, issue := range issues {
		if string(issue.Status) == status {
			filteredIssues = append(filteredIssues, issue)
		}
	}
	return filteredIssues
}
