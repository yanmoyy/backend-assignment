package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yanmoyy/backend-assignment/internal/api"
	"github.com/yanmoyy/backend-assignment/internal/client"
)

const baseURL = "http://localhost:8080"

func TestCreateIssue(t *testing.T) {
	cl := client.NewClient(baseURL)
	err := cl.Reset()
	if err != nil {
		t.Fatal(err)
	}
	id := uint(1)
	issue, err := cl.CreateIssue(
		client.CreateIssueParams{
			Title:       "버그 수정 필요",
			Description: "로그인 페이지에서 오류 발생",
			UserId:      &id,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, uint(1), issue.ID)
	assert.Equal(t, "버그 수정 필요", issue.Title)
	assert.Equal(t, "로그인 페이지에서 오류 발생", issue.Description)
	assert.Equal(t, api.StatusInProgress, issue.Status)
	assert.Equal(t, id, issue.User.ID)
	assert.Equal(t, "김개발", issue.User.Name)
}
