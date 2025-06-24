package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yanmoyy/backend-assignment/internal/api"
	"github.com/yanmoyy/backend-assignment/internal/client"
)

func TestUpdateIssue(t *testing.T) {
	cl := client.NewClient(baseURL)
	setupIssues(t, cl)
	t.Run("update title description", func(t *testing.T) {
		issue, err := cl.UpdateIssue(
			1,
			api.UpdateIssueParams{
				Title:       "new title",
				Description: "new description",
			},
		)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "new title", issue.Title)
		assert.Equal(t, "new description", issue.Description)
	})
	t.Run("set userId", func(t *testing.T) {
		id := uint(2)
		issue, err := cl.UpdateIssue(
			1,
			api.UpdateIssueParams{
				Title:  "로그인 버그 수정",
				Status: api.StatusInProgress,
				UserId: &id,
			},
		)
		assert.NoError(t, err)
		assert.Equal(t, api.StatusInProgress, issue.Status)
		assert.Equal(t, "로그인 버그 수정", issue.Title)
		assert.Equal(t, id, issue.User.ID)
		assert.Equal(t, "이디자인", issue.User.Name)
	})
}
