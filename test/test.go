package test

import (
	"testing"

	"github.com/yanmoyy/backend-assignment/internal/api"
	"github.com/yanmoyy/backend-assignment/internal/client"
)

const baseURL = "http://localhost:8080"

func setupIssues(t *testing.T, cl *client.Client) {
	err := cl.Reset()
	if err != nil {
		t.Fatal(err)
	}
	// create some issues
	id := uint(1)
	_, err = cl.CreateIssue(
		api.CreateIssueParams{
			Title:       "title1",
			Description: "description1",
			UserId:      &id, // userId exists (Status: IN_PROGRESS)
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	_, err = cl.CreateIssue(
		api.CreateIssueParams{
			Title:       "title2",
			Description: "description2", // userId does not exist (Status: PENDING)
		},
	)
	if err != nil {
		t.Fatal(err)
	}
}
