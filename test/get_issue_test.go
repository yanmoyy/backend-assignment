package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yanmoyy/backend-assignment/internal/api"
	"github.com/yanmoyy/backend-assignment/internal/client"
)

func TestGetIssue(t *testing.T) {
	cl := client.NewClient(baseURL)
	setupIssues(t, cl)
	t.Run("id = 1 (김개발)", func(t *testing.T) {
		issue, err := cl.GetIssue(1)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, uint(1), issue.ID)
		assert.Equal(t, "title1", issue.Title)
		assert.Equal(t, "description1", issue.Description)
		assert.Equal(t, api.StatusInProgress, issue.Status)
		assert.Equal(t, uint(1), issue.User.ID)
		assert.Equal(t, "김개발", issue.User.Name)
	})
	t.Run("id = 2", func(t *testing.T) {
		issue, err := cl.GetIssue(2)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, uint(2), issue.ID)
		assert.Equal(t, "title2", issue.Title)
		assert.Equal(t, "description2", issue.Description)
		assert.Equal(t, api.StatusPending, issue.Status)
	})
}
